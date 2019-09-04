package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	puppy "github.com/anz-bank/go-course/09_json/nickolee/pkg/puppy"
	store "github.com/anz-bank/go-course/09_json/nickolee/pkg/puppy/store"
	"gopkg.in/alecthomas/kingpin.v2"
)

var out io.Writer = os.Stdout

var (
	args     = os.Args[1:]
	fileName = kingpin.Flag("data", "This flag supplies a puppy JSON file name").Short('d').ExistingFile()
)

func main() {
	// parse flags/check if user even provided any
	parseFlags()

	// read json from file
	fileData := LoadFile() // fileData var to store raw bytes read in from file

	// parse/unmarshal the json and we now have []puppy (Go objects)
	puppies := unmarshalPuppies(fileData)

	// instantiate new puppy store
	store := store.NewMapStore()

	// create some puppies loaded in from json file and store in store (no pun intended lol)
	for _, pup := range puppies {
		pup := pup // to avoid scopelint error: Using a reference for the variable on range scope `pup`
		createdPuppy, _ := store.CreatePuppy(&pup)
		fmt.Fprintf(out, "Puppy with ID %d has been created\n", createdPuppy)
		retrievedPuppy, _ := store.ReadPuppy(createdPuppy)
		fmt.Fprintln(out, "Retrieved puppy:", retrievedPuppy)
	}
}

// checking for valid flags. At this point not checking for valid args provided with flags
func parseFlags() {
	// parse cli flags and check for flag parse error
	// I believe this also binds the user input (via terminal) to the var fileName
	if _, parseError := kingpin.CommandLine.Parse(args); parseError != nil {
		panic(parseError)
	}
}

// Load file and check for valid files
func LoadFile() (fileData []byte) {
	// if data flag is supplied. Read data from file and check for file argument error
	f, _ := os.Open(*fileName) // At this point Go knows about the file but not what's in it
	defer f.Close()            // to make sure the file is closed after this function returns
	// it seems that kingpin.CommandLine.Parse above checks for a variety of errors already such as path
	// and empty argument etc.

	// reading what's in file. Raw as in it is a raw slice of bytes at this point
	fileData, _ = ioutil.ReadAll(f) //no need to check err as kingpin has checked it already
	return fileData
}

// converts stream of bytes in file (json string data) into Go puppy objects
func unmarshalPuppies(fileData []byte) []puppy.Puppy {
	puppies := []puppy.Puppy{}
	err := json.Unmarshal(fileData, &puppies)
	if err != nil {
		fmt.Printf("Could not unmarshal puppies, error: %v\n", err)
		panic(err)
	}
	return puppies
}
