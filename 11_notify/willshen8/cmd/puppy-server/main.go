package main

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/anz-bank/go-course/11_notify/willshen8/pkg/puppy"
	"gopkg.in/alecthomas/kingpin.v2"
)

const lostPuppyServer = "http://localhost:8888/api/lostpuppy/"

var out io.Writer = os.Stdout

var (
	args      = os.Args[1:]
	fileName  = kingpin.Flag("data", "Puppy Json File Name").Short('d').File()
	port      = kingpin.Flag("port", "Port of the server").Short('p').Default("9999").Int()
	storer    = kingpin.Flag("store", "Map/Sync Store").Short('s').Default("map").Enum("map", "sync")
	lostPuppy = kingpin.Flag("lostPuppy", "server addr").Short('l').Default(lostPuppyServer).String()
)

// Config contains arguments parsed from commandline flags.
type Config struct {
	file      *os.File
	port      int
	storer    string
	lostPuppy string
}

func main() {
	cmdConfigs, err := parseCmdArgs(args)
	if err != nil {
		panic(err)
	}
	newStorer, storerErr := createStorer(cmdConfigs.storer)
	if storerErr != nil {
		panic(storerErr)
	}
	decodedPuppies, parseErr := parseJSONPuppies(cmdConfigs.file)
	if parseErr != nil {
		panic(parseErr)
	}
	portFlag, portErr := validatePortFlag(cmdConfigs)
	if portErr != nil {
		panic(portErr)
	}

	puppyStorer, _ := createPuppies(newStorer, decodedPuppies)
	puppyHandler := puppy.NewRestHandler(puppyStorer, cmdConfigs.lostPuppy)
	router := puppy.SetupRouter()
	puppyHandler.SetupRoutes(router)
	log.Fatal(http.ListenAndServe(":"+portFlag, router))
}

// creatStorer will take the commandline arg and create the appropriate store
func createStorer(store string) (puppy.Storer, error) {
	switch store {
	case "map":
		return puppy.NewMapStore(), nil
	case "sync":
		return puppy.NewSyncStore(), nil
	}
	return nil, errors.New("map/sync are the only acceptable flag values")
}

func parseCmdArgs(args []string) (Config, error) {
	var parsedConig Config
	if _, parseError := kingpin.CommandLine.Parse(args); parseError != nil {
		return parsedConig, parseError
	}
	return Config{file: *fileName, port: *port, storer: *storer, lostPuppy: *lostPuppy}, nil
}

// parseJSONPuppies read from io input and unmarshal them into puppy JSONs
func parseJSONPuppies(r io.Reader) ([]puppy.Puppy, error) {
	var puppies []puppy.Puppy
	jsonData, _ := ioutil.ReadAll(r) // error already handled by kingpin
	if err := json.Unmarshal(jsonData, &puppies); err != nil {
		return nil, err
	}
	return puppies, nil
}

// validatePortFlag check for valid port number entered from commandLine
func validatePortFlag(c Config) (string, error) {
	if c.port < 0 || c.port > 65535 {
		return "0", errors.New("invalid port number entered")
	}
	return strconv.Itoa(c.port), nil
}

// createPuppies takes an array of puppies and saves it to a puppy storer
func createPuppies(s puppy.Storer, puppies []puppy.Puppy) (puppy.Storer, error) {
	for _, jsonPuppy := range puppies {
		jsonPuppy := jsonPuppy
		if _, saveErr := s.CreatePuppy(&jsonPuppy); saveErr != nil {
			return nil, saveErr
		}
	}
	return s, nil
}
