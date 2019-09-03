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
	fileName = kingpin.Flag("data", "This flag supplies a puppy JSON file name").Short('d').String()
)

func main() {
	// parse flags and read json from file
	fileData := parseFlagsAndLoadFile() // fileData var to store raw bytes read in from file

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

// parse flags and load file
func parseFlagsAndLoadFile() (fileData []byte) {
	// parse cli flags and check for flag parse error
	if _, parseError := kingpin.CommandLine.Parse(args); parseError != nil {
		printUsage()
		panic(parseError) // this also includes a panic
	}

	// if data flag is supplied. Reading data from file and check for file argument error
	f, err := os.Open(*fileName) // At this point Go knows about the file but not what's in it
	if err != nil {
		printUsage()
		panic(err)
	}

	// reading what's in file. Raw as in it is a raw slice of bytes at this point
	fileData, _ = ioutil.ReadAll(f) // no need to check file error as kingpin has checked it already
	return fileData
}

// show usage guide if user is not providing args or providing the wrong args. Only applies to args not flags
func printUsage() {
	fmt.Println("No/wrong data file and/or flag was provided. See below for usage guidelines: ")
	fmt.Println("The arguments you provided: ", args)
	fmt.Println("----------------------------------------------------------")
	fmt.Print(`usage: main [<flags>]

Flags:
	--help           Show context-sensitive help (also try --help-long and --help-man).
	-d, --data=DATA  This flag supplies a puppy JSON file name	
`)
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
