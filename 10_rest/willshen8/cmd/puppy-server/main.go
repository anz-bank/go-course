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
	fileName = kingpin.Flag("data", "Puppy Json File Name").Short('d').ExistingFile()
	port     = kingpin.Flag("port", "Port of the server").Short('p').Default("8888").Int()
	storer   = kingpin.Flag("store", "Map/Sync Store").Short('s').Default("map").String()
)

// Config contains arguments parsed from commandline flags.
type Config struct {
	file   string
	port   int
	storer string
}

func main() {
	cmdConfigs := parseCmdArgs(args)
	newStorer, _ := createStorer(cmdConfigs)
	decodedPuppies, _ := parseFileFlag(cmdConfigs)
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
func createStorer(c *Config) (puppy.Storer, error) {
	switch c.storer {
	case "map":
		return puppy.NewMapStore(), nil
	case "sync":
		return puppy.NewSyncStore(), nil
	}
	return nil, errors.New("map/sync are the only acceptable flag values")
}

func parseCmdArgs(args []string) *Config {
	if _, parseError := kingpin.CommandLine.Parse(args); parseError != nil {
		panic(parseError)
	}
	return &Config{file: *fileName, port: *port, storer: *storer}
}

func parseFileFlag(c *Config) ([]puppy.Puppy, error) {
	var file *string
	var puppies []puppy.Puppy
	file = &c.file
	jsonData := readFile(*file)

	if err := json.Unmarshal(jsonData, &puppies); err != nil {
		return nil, err
	}
	return puppies, nil
}

//parsePortFlag check for valid port number entered from commandLine
func parsePortFlag(c *Config) (string, error) {
	if c.port < 0 || c.port > 65535 {
		return "8888", errors.New("invalid port number entered, default port 8888 will be used")
	}
	return strconv.Itoa(c.port), nil
}

func readFile(file string) []byte {
	jsonData, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return jsonData
}

// createPuppies takes an array of puppies and saves it to a puppy storer
func createPuppies(s puppy.Storer, decodedPuppies []puppy.Puppy) (puppy.Storer, error) {
	for _, jsonPuppy := range decodedPuppies {
		jsonPuppy := jsonPuppy //fixes linter issue
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
