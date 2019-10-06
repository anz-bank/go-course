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

	"github.com/anz-bank/go-course/10_rest/nickolee/pkg/puppy/store"

	puppy "github.com/anz-bank/go-course/10_rest/nickolee/pkg/puppy"
	rest "github.com/anz-bank/go-course/10_rest/nickolee/pkg/puppy/rest"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"gopkg.in/alecthomas/kingpin.v2"
)

var out io.Writer = os.Stdout

var (
	args = os.Args[1:]
	// this loads your file once parseFlags() is run
	// these are methods too prepare a variable. Basically telling kingpin how to read the flag from cli
	fileName = kingpin.Flag("data", "This flag supplies a puppy JSON file name").
			Short('d').String()
	port   = kingpin.Flag("port", "Port to listen on and serve").Short('p').Default("7777").Int()
	storer = kingpin.Flag("store", "Map/Sync Store").Short('s').Default("map").Enum("map", "sync")
)

// Config contains arguments parsed from commandline flags
type Config struct {
	file   *string
	port   int
	storer string
}

func main() {
	// parse flags/check if user even provided any
	cliConfig, err := parseFlags(args)
	if err != nil {
		fmt.Println(err)
	}

	// instantiate new PuppyHandlerAndStorer based on user specified value of storer
	// create puppy.Storer
	newStorer, _ := createStore(cliConfig.storer) // no need to check err as it defaults to "map"

	// create PuppyHandlerAndStorer - the Rest API wrapper around puppy.Storer
	phs := rest.NewPuppyHandlerAndStorer(newStorer)

	// read JSON from file + parse/unmarshal the json and we now have []puppy (Go objects)
	puppies, err := readFileAndUnmarshalPuppies(cliConfig.file)
	if err != nil {
		panic(err)
	}

	// create some puppies loaded in from json file and store in store (no pun intended lol)
	for _, pup := range puppies {
		pup := pup // to avoid scopelint error: Using a reference for the variable on range scope `pup`
		createdPuppy, _ := phs.Storage.CreatePuppy(&pup)
		fmt.Fprintf(out, "Puppy with ID %d has been created\n", createdPuppy)
	}

	// setting up REST API wrapper
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	rest.SetupRoutes(r, *phs)

	// get the user specified port
	addr := ":" + strconv.Itoa(cliConfig.port)

	fmt.Printf("Starting server on port %s.\n", addr)
	log.Print(http.ListenAndServe(addr, r))
}

// checking for valid flags. At this point not checking for valid args provided with flags
func parseFlags(a []string) (Config, error) {
	// parse cli flags and check for flag parse error
	// I believe this also binds the user input (via terminal) to the var fileName, port and storer
	// but now I am extracting them and storing in Config struct in an easy to consume form
	var parsedConfig Config
	if _, parseError := kingpin.CommandLine.Parse(a); parseError != nil {
		return parsedConfig, parseError
	}
	return Config{file: fileName, port: *port, storer: *storer}, nil
}

// Read data from file (converts stream of bytes in file into json string data)
// And also unmarshals JSON string into Go puppy objects
func readFileAndUnmarshalPuppies(f *string) ([]puppy.Puppy, error) {
	// handle if no file provided by user
	if f == nil {
		return []puppy.Puppy{}, nil
	}

	// open file
	file, err := os.Open(*f)
	if err != nil {
		return nil, err
	}

	puppies := []puppy.Puppy{}
	// rawJSONData is just a byte stream at this point
	rawJSONData, _ := ioutil.ReadAll(file) // error already handled by kingpin
	err = json.Unmarshal(rawJSONData, &puppies)
	if err != nil {
		return nil, err
	}
	return puppies, nil
}

// createStore will take the commandline arg and create the relevant store
func createStore(s string) (puppy.Storer, error) {
	switch s {
	case "map":
		return store.NewMapStore(), nil
	case "sync":
		return store.NewSyncStore(), nil
	}
	return nil, errors.New("map/sync are the only acceptable flag values")
}
