package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/anz-bank/go-course/09_json/kasunfdo/pkg/puppy"
	"github.com/anz-bank/go-course/09_json/kasunfdo/pkg/puppy/store"
)

var out io.Writer = os.Stdout

func main() {
	var dataFile *os.File
	kingpin.Flag("data", "path to puppy data file").Short('d').Default("data.json").FileVar(&dataFile)
	_, err := kingpin.CommandLine.Parse(os.Args[1:])
	if err != nil {
		panic(err)
	}

	if err := createPuppies(dataFile, store.NewMapStore()); err != nil {
		panic(err)
	}
}

func createPuppies(dataFile io.Reader, store puppy.Storer) error {
	data, err := ioutil.ReadAll(dataFile)
	if err != nil {
		return puppy.NewError(puppy.ErrInternal)
	}

	var puppies []puppy.Puppy
	if err = json.Unmarshal(data, &puppies); err != nil {
		return puppy.NewError(puppy.ErrBadFormat)
	}

	for _, p := range puppies {
		var id uint64
		if id, err = store.CreatePuppy(p); err != nil {
			return err
		}
		fmt.Fprintf(out, "%v (id: %d) added to store\n", p, id)
	}

	return nil
}
