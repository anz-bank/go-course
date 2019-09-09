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

	"github.com/go-chi/chi"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	args           = os.Args[1:]
	data           = kingpin.Flag("data", "data file").Short('d').Default("puppydata/data.json").ExistingFile()
	port           = kingpin.Flag("port", "PORT").Short('p').Default("8080").String()
	db             = kingpin.Flag("store", "STORE").Short('s').Default("sync").String()
	out  io.Writer = os.Stdout
)

func main() {
	_, err := kingpin.CommandLine.Parse(args)
	if err != nil {
		fmt.Fprintf(out, "Command line could not be parsed, error: %v", err)
		return
	}

	d, _ := ioutil.ReadFile(*data)
	puppies, err := unmarshalPuppies(d)
	if err != nil {
		panic(err)
	}

	store, err := newStore(*db)
	if err != nil {
		fmt.Fprintf(out, "Could not setup database, error: %v", err)
		return
	}
	err = load(store, puppies)
	if err != nil {
		fmt.Fprintf(out, "Could not load database, error: %v", err)
		return
	}

	fmt.Fprintf(out, "Store of puppies:%v", store)

	r := chi.NewRouter()

	rs := puppy.RestStorer{Db: store}
	r.Get(rs.GetPuppyRoute(), rs.GetPuppy)
	r.Post(rs.PostPuppyRoute(), rs.CreatePuppy)
	r.Put(rs.PutPuppyRoute(), rs.UpdatePuppy)
	r.Delete(rs.DeletePuppyRoute(), rs.DeletePuppy)

	portValue := fmt.Sprintf(":%v", *port)
	serv := http.Server{Addr: portValue, Handler: r}
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

func newStore(db string) (puppy.Storer, error) {
	switch db {
	case "map":
		return store.NewMapStore(), nil
	case "sync":
		return store.NewSyncStore(), nil
	}
	return nil, fmt.Errorf("invalid storage")
}

func unmarshalPuppies(d []byte) ([]puppy.Puppy, error) {
	puppies := []puppy.Puppy{}
	if err := json.Unmarshal(d, &puppies); err != nil {
		return nil, err
	}
	return puppies, nil
}
