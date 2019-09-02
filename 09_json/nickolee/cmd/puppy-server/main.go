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
	args        = os.Args[1:]
	fileName    = kingpin.Flag("data", "This flag supplies a puppy JSON file name").Short('d').String()
	fileData    []byte
	unmarshaled puppy.Puppy
)

func main() {
	// by this point we have read from a file and parsed/unmarshaled the json and we now have []puppy (Go objects)
	puppies, _ := unmarshalPuppies(fileData)
	fmt.Printf("%#v \n", puppies)
	fmt.Println("Number of puppies loaded from file: ", len(puppies))
	// fmt.Printf("%v \n", puppies)

	// instantiate new puppy store
	store := store.NewMapStore()

	// create some puppies loaded in from json file and store in store (no pun intended lol)
	for _, pup := range puppies {
		createdPuppy, _ := store.CreatePuppy(&pup)
		fmt.Fprintf(out, "Puppy with ID %d has been created\n", createdPuppy)

		retrievedPuppy, _ := store.ReadPuppy(createdPuppy)
		fmt.Fprintln(out, "Retrieved puppy:", retrievedPuppy)
	}

	// updateResult := store.UpdatePuppy(createdPuppy, &puppy.Puppy{Breed: "Arcanine", Colour: "Level 100", Value: 9300.90})
	// fmt.Fprintln(out, "Update puppy operation result:", updateResult)

	// deleteResult := store.DeletePuppy(createdPuppy)
	// fmt.Fprintln(out, "Delete puppy operation result:", deleteResult)
}

// this runs first
func init() {
	// parse flags
	kingpin.Parse()

	// if data flag is supplied. Reading data from file
	if *fileName != "" {
		// At this point Go knows about the file but not what's in it
		f, err := os.Open(*fileName)
		if err != nil {
			panic(err)
		}

		// reading what's in file. Raw as in it is a raw slice of bytes at this point
		fileData, err = ioutil.ReadAll(f)
		if err != nil {
			panic(err)
		}
	} else {
		// show usage guide if user is not providing args or providing the wrong flags
		fmt.Println("No data file was provided. See below for usage guidelines: ")
		fmt.Println("The arguments you provided: ", os.Args[1:])
		fmt.Println("----------------------------------------------------------")
		printUsage()
	}
}

func printUsage() {
	kingpin.Usage()
	os.Exit(1)
}

// converts stream of bytes in file (json string data) into Go puppy objects
func unmarshalPuppies(data []byte) ([]puppy.Puppy, error) {
	puppies := []puppy.Puppy{}
	err := json.Unmarshal(fileData, &puppies)
	if err != nil {
		fmt.Printf("Could not unmarshall puppies, error: %v", err)
		return nil, err
	}
	return puppies, nil

	// fmt.Println("unmarshaled - so it is back to a Go object now: ", unmarshaled)
	// fmt.Printf("unmarshaled - so it is back to a Go object now: %#v\n", unmarshaled)
}
