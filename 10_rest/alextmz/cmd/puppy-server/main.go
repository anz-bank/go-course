package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/anz-bank/go-course/10_rest/alextmz/pkg/puppy"
	"github.com/anz-bank/go-course/10_rest/alextmz/pkg/puppy/store"
	"github.com/go-chi/chi"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	args = os.Args[1:]
	// Default("./test/valid-formatted-json.json")
	flagfile  = kingpin.Flag("data", "JSON file to read").Short('d').Required().String()
	flagport  = kingpin.Flag("port", "TCP port to listen on").Short('p').Required().Uint16()
	flagstore = kingpin.Flag("store", "Backing store to use").Short('s').Required().Enum("map", "sync")
)

var out io.Writer = os.Stdout

func main() {

	_, err := kingpin.CommandLine.Parse(args)
	if err != nil {
		fmt.Fprintf(out, "error %s", err.Error())
		kingpin.Usage()
		return
	}

	// var err error
	// defer func() {
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// }()

	_, err = kingpin.CommandLine.Parse(args)
	if err != nil {
		// fmt.Fprintf(out, "error!")
		// return
		// log.Fatalf("Failed to parse args.\n%v", err)
		kingpin.FatalUsage("failed to parse arguments: %s, received arguments: %#v", err.Error(), args)
		// return
	}

	puppies, err := readfile(*flagfile, false)
	if err != nil {
		fmt.Fprintf(out, "failed to read file %s: %v\n", *flagfile, err)
		return
	}

	var puppystore puppy.Storer
	switch *flagstore {
	case "map":
		puppystore = store.NewmapStore()
	default:
		puppystore = store.NewSyncStore()
	}

	n, err := storepuppies(puppystore, puppies)
	if err != nil {
		fmt.Fprintf(out, "%v\n", err)
		return
	}
	fmt.Fprintf(out, "Starting puppyserver with options:\n")
	fmt.Fprintf(out, "file   = %s\nport  = %d\nstore = %s\n", *flagfile, *flagport, *flagstore)
	fmt.Fprintf(out, "Loaded %d puppies.\n", n)
	//printpuppies(puppystore, n)
	h := puppy.HTTPHandler{Store: puppystore}
	r := chi.NewRouter()

	puppy.SetupRoutes(r, h)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(int(*flagport)), r))
}

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

func storepuppies(store puppy.Storer, puppies []puppy.Puppy) (int, error) {
	n := 0
	for _, v1 := range puppies {
		n++
		v := v1
		err := store.CreatePuppy(&v)
		if err != nil {
			return 0, err
		}
	}
	return n, nil
}

func readfile(filename string, testing bool) ([]puppy.Puppy, error) {
	jsonfile, err := os.Open(filename)
	if err != nil {
		return []puppy.Puppy{}, err
	}
	defer jsonfile.Close()
	bytes, err := ioutil.ReadAll(jsonfile)
	if testing {
		err = errors.New("mock error")
	}
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
