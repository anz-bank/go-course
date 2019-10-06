package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	puppy "github.com/anz-bank/go-course/10_rest/runnerdave/pkg/puppy"
	store "github.com/anz-bank/go-course/10_rest/runnerdave/pkg/puppy/store"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	args           = os.Args[1:]
	data           = kingpin.Flag("data", "data file").Short('d').Default("puppydata/data.json").ExistingFile()
	port           = kingpin.Flag("port", "PORT").Short('p').Default("8080").Int()
	db             = kingpin.Flag("store", "STORE").Short('s').Default("sync").Enum("sync", "map")
	out  io.Writer = os.Stdout
	rs   puppy.RestServer
)

func main() {
	_, err := kingpin.CommandLine.Parse(args)
	if err != nil {
		panic(err)
	}

	d, _ := ioutil.ReadFile(*data)
	puppies, err := unmarshalPuppies(d)
	if err != nil {
		panic(err)
	}

	s := newStore(*db)
	err = load(s, puppies)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(out, "Store of puppies:%v", s)

	handler := rs.SetupRoutes(s)

	portValue := fmt.Sprintf(":%v", *port)
	serv := http.Server{Addr: portValue, Handler: handler}
	serr := serv.ListenAndServe()
	if serr != http.ErrServerClosed {
		fmt.Fprintf(out, "Could not start puppy server: %v", serr)
	}
}

func load(s puppy.Storer, puppies []puppy.Puppy) error {
	for _, p := range puppies {
		err := s.CreatePuppy(p)
		if err != nil {
			return err
		}
	}
	return nil
}

func newStore(db string) puppy.Storer {
	switch db {
	case "map":
		return store.NewMapStore()
	case "sync":
		return store.NewSyncStore()
	}
	return nil
}

func unmarshalPuppies(d []byte) ([]puppy.Puppy, error) {
	puppies := []puppy.Puppy{}
	if err := json.Unmarshal(d, &puppies); err != nil {
		return nil, err
	}
	return puppies, nil
}
