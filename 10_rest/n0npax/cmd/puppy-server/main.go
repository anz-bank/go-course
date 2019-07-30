package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"

	puppy "github.com/anz-bank/go-course/10_rest/n0npax/pkg/puppy"
	store "github.com/anz-bank/go-course/10_rest/n0npax/pkg/puppy/store"
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
	s, err := createStorer(c)
	if err != nil {
		return err
	}
	if err := feedStorer(*c, s); err != nil {
		return err
	}
	return puppy.RestBackend(s).Run(fmt.Sprintf(":%d", c.port))
}

type config struct {
	puppyFile io.Reader
	sType     string
	port      int
}

func parseArgs(args []string) (config, error) {
	var storeType string
	var port int
	var puppyFile *os.File
	kingpin.Flag("data", "path to file with puppies data").Short('d').FileVar(&puppyFile)
	kingpin.Flag("port", "Port number").Short('p').Default("8181").IntVar(&port)
	kingpin.Flag("store", "Store type").Short('s').Default("map").EnumVar(&storeType, "map", "sync")
	_, err := kingpin.CommandLine.Parse(args)
	return config{puppyFile, storeType, port}, err
}

func createStorer(c *config) (puppy.Storer, error) {
	switch c.sType {
	case "sync":
		return store.NewSyncStore(), nil
	case "map":
		return store.NewMemStore(), nil
	default:
		return nil, errors.New("unknown storer type")
	}
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

func feedStorer(c config, s puppy.Storer) error {
	puppies, err := readPuppies(c.puppyFile)
	if err != nil {
		return err
	}
	for _, p := range puppies {
		p := p
		if _, err := s.CreatePuppy(&p); err != nil {
			return err
		}
	}
	return nil
}
