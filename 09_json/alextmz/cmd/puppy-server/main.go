package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/anz-bank/go-course/09_json/alextmz/pkg/puppy"
	"github.com/anz-bank/go-course/09_json/alextmz/pkg/puppy/store"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	args = os.Args[1:]
	file = kingpin.Flag("data", "JSON file to read").Short('d').String()
)

var out io.Writer = os.Stdout

func main() {
	_, err := kingpin.CommandLine.Parse(args)
	if err != nil {
		fmt.Fprintf(out, "error parsing command line: %v\n", err)
		return
	}

	jsonfile, err := os.Open(*file)
	if err != nil {
		fmt.Fprintf(out, "error opening file: %v\n", err)
		return
	}
	defer jsonfile.Close()

	puppies, err := readfile(jsonfile)
	if err != nil {
		fmt.Fprintf(out, "error reading JSON file: %v\n", err)
		return
	}

	puppystore := store.NewMapStore()

	n, err := storepuppies(puppystore, puppies)
	if err != nil {
		fmt.Fprintf(out, "error storing puppies: %v\n", err)
		return
	}

	printpuppies(puppystore, n)
}

// printpuppies print n puppies contained in the store s
func printpuppies(s puppy.Storer, n int) {
	for i := 1; i <= n; i++ {
		p, err := s.ReadPuppy(i)
		if err != nil {
			fmt.Fprintf(out, "%v\n", err)
			return
		}

		fmt.Fprintf(out, "Printing puppy id %d: %#v\n", i, p)
	}
}

// storepuppies store all puppies contained in slice 'puppies'
// into the store 'store', returning either (number of puppies stored, nil)
// if there is no error or (0, error) if there was an error.
func storepuppies(store puppy.Storer, puppies []puppy.Puppy) (int, error) {
	n := 0

	for _, v1 := range puppies {
		v := v1

		err := store.CreatePuppy(&v)
		if err != nil {
			return 0, err
		}
		n++
	}

	return n, nil
}

func readfile(file io.Reader) ([]puppy.Puppy, error) {
	bytes, err := ioutil.ReadAll(file)

	if err != nil {
		return []puppy.Puppy{}, err
	}

	var puppies []puppy.Puppy

	err = json.Unmarshal(bytes, &puppies)
	if err != nil {
		return []puppy.Puppy{}, err
	}

	return puppies, nil
}
