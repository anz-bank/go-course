package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/url"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"

	puppy "github.com/anz-bank/go-course/11_notify/n0npax/pkg/puppy"
	store "github.com/anz-bank/go-course/11_notify/n0npax/pkg/puppy/store"
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

var runPuppyServer = func(c *config) error {
	s, err := createStorer(c)
	if err != nil {
		return err
	}
	if err := feedStorer(c.puppyReader, s); err != nil {
		return err
	}
	puppy.LostPuppyURL = c.lostpuppyURL
	return puppy.RestBackend(s).Run(fmt.Sprintf(":%d", c.port))
}

type config struct {
	puppyReader  io.Reader
	sType        string
	port         int
	lostpuppyURL string
}

func parseArgs(args []string) (config, error) {
	var storeType string
	var lostpuppyURL *url.URL
	var port int
	var puppyFile *os.File
	kingpin.Flag("data", "path to file with puppies data").Short('d').FileVar(&puppyFile)
	kingpin.Flag("port", "Port number").Short('p').Default("8181").IntVar(&port)
	kingpin.Flag("lostpuppy", "lostpuppy service endpoint").Short('l').Required().URLVar(&lostpuppyURL)
	kingpin.Flag("store", "Store type").Short('s').Default("map").EnumVar(&storeType, "map", "sync")
	_, err := kingpin.CommandLine.Parse(args)
	return config{puppyFile, storeType, port, lostpuppyURL.String()}, err
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
		return nil, errors.New("failed to read puppies from file")
	}
	var puppies []puppy.Puppy
	if err = json.Unmarshal(b, &puppies); err != nil {
		return nil, err
	}
	return puppies, nil
}

func feedStorer(r io.Reader, s puppy.Storer) error {
	puppies, err := readPuppies(r)
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
