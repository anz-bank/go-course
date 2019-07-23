package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/anz-bank/go-course/09_json/willshen8/pkg/puppy"
	store "github.com/anz-bank/go-course/09_json/willshen8/pkg/puppy/store"
	"gopkg.in/alecthomas/kingpin.v2"
)

var out io.Writer = os.Stdout

var (
	args     = os.Args[1:]
	fileName = kingpin.Flag("data", "Puppy Json File Name").Short('d').String()
)

func main() {
	var file *string
	var jsonData []byte
	var decodedPuppies []puppy.Puppy
	var mystore = store.NewMapStore()

	file = processCommandLineArgs(args)
	jsonData = readFile(*file)

	if err := json.Unmarshal(jsonData, &decodedPuppies); err != nil {
		panic(err)
	}

	if savePuppyErr := createPuppies(mystore, decodedPuppies); savePuppyErr != nil {
		panic(savePuppyErr)
	}
}

func processCommandLineArgs(args []string) *string {
	if _, parseError := kingpin.CommandLine.Parse(args); parseError != nil {
		panic(parseError)
	}
	return fileName
}

func readFile(file string) []byte {
	jsonData, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return jsonData
}

func createPuppies(store store.Storer, decodedPuppies []puppy.Puppy) error {
	for _, jsonPuppy := range decodedPuppies {
		jsonPuppy := jsonPuppy //fixes linter issue
		fmt.Fprintln(out, jsonPuppy)
		if _, saveErr := store.CreatePuppy(&jsonPuppy); saveErr != nil {
			return saveErr
		}
	}
	return nil
}
