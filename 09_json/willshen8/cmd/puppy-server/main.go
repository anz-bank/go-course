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
	fileName = kingpin.Flag("data", "Puppy Json File Name").Short('d').ExistingFile()
)

func main() {
	mystore := store.NewMapStore()
	file := parseCmdArgs(args)
	decodedPuppies := readFileAndMarshal(*file)

	if savePuppyErr := createPuppies(mystore, decodedPuppies); savePuppyErr != nil {
		panic(savePuppyErr)
	}
}

func parseCmdArgs(args []string) *string {
	if _, parseError := kingpin.CommandLine.Parse(args); parseError != nil {
		panic(parseError)
	}
	return fileName
}

func readFileAndMarshal(file string) []puppy.Puppy {
	jsonData, _ := ioutil.ReadFile(file) //no need to check file error as kingpin check it already
	var decodedPuppies []puppy.Puppy
	if err := json.Unmarshal(jsonData, &decodedPuppies); err != nil {
		panic(err)
	}
	return decodedPuppies
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
