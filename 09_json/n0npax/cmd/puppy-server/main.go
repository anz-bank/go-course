package main

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"

	puppy "github.com/anz-bank/go-course/09_json/n0npax/pkg/puppy"
	store "github.com/anz-bank/go-course/09_json/n0npax/pkg/puppy/store"
)

var (
	logFatalf = log.Fatal
	parser    = parseArgs
)

func main() {
	config, err := parser(os.Args[1:])
	if err != nil {
		logFatalf(err)
	}
	logFatalf(runPuppyServer(&config))
}

func runPuppyServer(c *config) error {
	if err := createStorer(c); err != nil {
		return err
	}
	if err := feedStorer(*c); err != nil {
		return err
	}
	return nil
}

func parseArgs(args []string) (config, error) {
	var storeType string
	var puppyFile *os.File
	kingpin.Flag("data", "path to file with puppies data").Short('d').FileVar(&puppyFile)
	kingpin.Flag("store", "Store type").Short('s').Default("map").EnumVar(&storeType, "map", "sync")
	_, err := kingpin.CommandLine.Parse(args)
	return config{puppyFile, storeType, nil}, err
}

func createStorer(c *config) error {
	switch c.sType {
	case "sync":
		c.storer = store.NewSyncStore()
	case "map":
		c.storer = store.NewMemStore()
	default:
		return errors.New("unknown storer type")
	}
	return nil
}

func readPuppies(r io.Reader) ([]puppy.Puppy, error) {
	if r == (*os.File)(nil) {
		return []puppy.Puppy{}, nil
	}
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, errors.New("error during reading puppies from file")
	}
	var puppies []puppy.Puppy
	if err = json.Unmarshal(b, &puppies); err != nil {
		return nil, errors.New(string(b))
	}
	return puppies, nil
}

func feedStorer(c config) error {
	puppies, err := readPuppies(c.puppyFile)
	if err != nil {
		return err
	}
	for _, p := range puppies {
		p := p
		if _, err := c.storer.CreatePuppy(&p); err != nil {
			return err
		}
	}
	return nil
}

type config struct {
	puppyFile io.Reader
	sType     string
	storer    puppy.Storer
}
