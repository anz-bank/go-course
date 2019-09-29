package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/anz-bank/go-course/10_rest/alextmz/pkg/puppy"
	"github.com/anz-bank/go-course/10_rest/alextmz/pkg/puppy/store"
	"github.com/anz-bank/go-course/10_rest/alextmz/pkg/rest"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	args      = os.Args[1:]
	flagfile  = kingpin.Flag("data", "JSON file to read").Short('d').Default("./test/valid-formatted-json.json").String()
	flagport  = kingpin.Flag("port", "TCP port to listen on").Short('p').Default("7735").Uint16()
	flagstore = kingpin.Flag("store", "Backing store to use").Short('s').Default("sync").Enum("map", "sync")

	out io.Writer = os.Stdout

	// shutdownhttp signals main() to... shutdown the http server
	shutdownhttp = make(chan bool)

	// syncoutput signals whoever interested that main() output is done
	syncoutput = make(chan bool, 1)
)

func main() {
	if _, err := kingpin.CommandLine.Parse(args); err != nil {
		fmt.Fprintf(out, "error parsing command line: %v\n", err)
		return
	}

	jsonfile, err := os.Open(*flagfile)
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

	var puppystore puppy.Storer

	switch *flagstore {
	case "map":
		puppystore = store.NewMapStore()
	case "sync":
		puppystore = store.NewSyncStore()
	}

	n, err := storepuppies(puppystore, puppies)
	if err != nil {
		fmt.Fprintf(out, "error storing puppies: %v\n", err)
		return
	}

	fmt.Fprintf(out, "Starting puppyserver with options:\n")
	fmt.Fprintf(out, "file   = %s\nport  = %d\nstore = %s\n", *flagfile, *flagport, *flagstore)
	fmt.Fprintf(out, "Loaded %d puppies.\n", n)
	printpuppies(puppystore, n)

	h := rest.HTTPHandler{Store: puppystore}
	s := http.Server{
		Addr:    ":" + strconv.Itoa(int(*flagport)),
		Handler: h,
	}

	// synchttpshutdown forces main() to wait for http.Shutdown to complete.
	// not necessarily needed here, but documented as good practice, so I used it.
	synchttpshutdown := make(chan bool)

	// this goroutine just waits blocked for something in the shutdownhttp channel.
	// if it gets anything, signals the server to stop gracefully.
	go func() {
		<-shutdownhttp
		_ = s.Shutdown(context.Background())
		close(synchttpshutdown)
	}()

	// signals whoever is listening that there is no more
	// io.Writer output to be done from main().
	syncoutput <- true

	err = s.ListenAndServe()
	if err != nil && err == http.ErrServerClosed {
		<-synchttpshutdown
	}
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
	for _, v := range puppies {
		v := v

		err := store.CreatePuppy(&v)
		if err != nil {
			return 0, err
		}
	}

	return len(puppies), nil
}

func readfile(file io.Reader) ([]puppy.Puppy, error) {
	bytes, err := ioutil.ReadAll(file)

	if err != nil {
		return []puppy.Puppy{}, err
	}

	var puppies []puppy.Puppy

	if err = json.Unmarshal(bytes, &puppies); err != nil {
		return []puppy.Puppy{}, err
	}

	return puppies, nil
}
