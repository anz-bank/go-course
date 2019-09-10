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
			Short('d').Default("../../puppy-data/puppies.json").File()
	port   = kingpin.Flag("port", "Port to listen on and serve").Short('p').Default("7777").Int()
	storer = kingpin.Flag("store", "Map/Sync Store").Short('s').Default("map").Enum("map", "sync")
)

// Config contains arguments parsed from commandline flags
type Config struct {
	file   *os.File
	port   int
	storer string
}

func main() {
	// parse flags/check if user even provided any
	cliConfig, err := parseFlags(args)
	if err != nil {
		fmt.Println(err)
	}

	// read JSON from file + parse/unmarshal the json and we now have []puppy (Go objects)
	puppies, _ := readFileAndUnmarshalPuppies(cliConfig.file)
	if err != nil {
		panic(err)
	}

	// instantiate new PuppyHandlerAndStorer based on user specified value of storer
	// create puppy.Storer
	newStorer, _ := createStore(cliConfig.storer) // no need to check err as it defaults to "map"

	// create PuppyHandlerAndStorer - the Rest API wrapper around puppy.Storer
	phs := rest.NewPuppyHandlerAndStorer(newStorer)

	// get the user specified port
	addr, err := parsePortFlag(cliConfig.port)
	if err != nil {
		panic(err)
	}

	// create some puppies loaded in from json file and store in store (no pun intended lol)
	for _, pup := range puppies {
		pup := pup // to avoid scopelint error: Using a reference for the variable on range scope `pup`
		createdPuppy, _ := phs.Storage.CreatePuppy(&pup)
		fmt.Fprintf(out, "Puppy with ID %d has been created\n", createdPuppy)
		retrievedPuppy, _ := phs.Storage.ReadPuppy(createdPuppy)
		fmt.Fprintln(out, "Retrieved puppy:", retrievedPuppy)
	}

	// setting up REST API wrapper
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	rest.SetupRoutes(r, *phs)

	fmt.Printf("Starting server on port %s. Try:\n", addr)
	fmt.Printf("  curl localhost%s/api/puppy -d ", addr)
	fmt.Println(`'{"breed": "Snorlax", "colour": "Blue-ish", "value": 8888}'`)
	fmt.Printf("  curl localhost%s/api/puppy/4 \n", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Println("Error: ", err)
	}
}

// checking for valid flags. At this point not checking for valid args provided with flags
func parseFlags(args []string) (Config, error) {
	// parse cli flags and check for flag parse error
	// I believe this also binds the user input (via terminal) to the var fileName, port and storer
	// but now I am extracting them and storing in Config struct in an easy to consume form
	var parsedConfig Config
	if _, parseError := kingpin.CommandLine.Parse(args); parseError != nil {
		return parsedConfig, parseError
	}
	return Config{file: *fileName, port: *port, storer: *storer}, nil
}

// Read data from file (converts stream of bytes in file into json string data)
// And also unmarshals JSON string into Go puppy objects
func readFileAndUnmarshalPuppies(r io.Reader) ([]puppy.Puppy, error) {
	puppies := []puppy.Puppy{}
	// rawJSONData is just a byte stream at this point
	rawJSONData, _ := ioutil.ReadAll(r) // error already handled by kingpin
	err := json.Unmarshal(rawJSONData, &puppies)
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

// parsePortFlag checks for valid port number entered from commandLine
func parsePortFlag(p int) (string, error) {
	if p < 0 || p > 65535 {
		return "0", errors.New("invalid port number entered")
	}
	return ":" + strconv.Itoa(p), nil
}
