package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/anz-bank/go-course/10_rest/willshen8/pkg/puppy"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"gopkg.in/alecthomas/kingpin.v2"
)

var out io.Writer = os.Stdout

var (
	args     = os.Args[1:]
	fileName = kingpin.Flag("data", "Puppy Json File Name").Short('d').File()
	port     = kingpin.Flag("port", "Port of the server").Short('p').Default("8888").Int()
	storer   = kingpin.Flag("store", "Map/Sync Store").Short('s').Default("map").Enum("map", "sync")
)

// Config contains arguments parsed from commandline flags.
type Config struct {
	file   *os.File
	port   int
	storer string
}

func main() {
	cmdConfigs, err := parseCmdArgs(args)
	if err != nil {
		fmt.Println(err)
	}
	newStorer, storerErr := createStorer(cmdConfigs.storer)
	if storerErr != nil {
		panic(storerErr)
	}
	decodedPuppies, parseErr := parseJSONPuppies(cmdConfigs.file)
	if parseErr != nil {
		panic(parseErr)
	}
	portFlag, portErr := parsePortFlag(cmdConfigs)
	if portErr != nil {
		fmt.Println(portErr)
	}
	port := ":" + portFlag

	puppyStorer, _ := createPuppies(newStorer, decodedPuppies)
	fmt.Fprintln(out, puppyStorer)
	puppyHandler := puppy.NewRestHandler(puppyStorer)
	router := SetupRouter()
	puppy.SetupRoutes(router, *puppyHandler)
	log.Fatal(http.ListenAndServe(port, router))
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
	return Config{file: *fileName, port: *port, storer: *storer}, nil
}

// parseJSONPuppies read from io input and unmarshal them into puyppy JSONS
func parseJSONPuppies(r io.Reader) ([]puppy.Puppy, error) {
	var puppies []puppy.Puppy
	jsonData, _ := ioutil.ReadAll(r) // error already handled by kingpin
	if err := json.Unmarshal(jsonData, &puppies); err != nil {
		return nil, err
	}
	return puppies, nil
}

// parsePortFlag check for valid port number entered from commandLine
func parsePortFlag(c Config) (string, error) {
	if c.port < 0 || c.port > 65535 {
		return "0", errors.New("invalid port number entered")
	}
	return strconv.Itoa(c.port), nil
}

// createPuppies takes an array of puppies and saves it to a puppy storer
func createPuppies(s puppy.Storer, puppies []puppy.Puppy) (puppy.Storer, error) {
	for _, jsonPuppy := range puppies {
		jsonPuppy := jsonPuppy // fixes linter issue
		if _, saveErr := s.CreatePuppy(&jsonPuppy); saveErr != nil {
			return nil, saveErr
		}
	}
	return s, nil
}

// SetupRouter takes the port number parsed from cmd and starts the server.
func SetupRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	return r
}
