package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

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
	test           = kingpin.Flag("test", "TEST").Short('t').Default("false").String()
	out  io.Writer = os.Stdout
	s    puppy.Storer
	c    = make(chan int)
)

func main() {
	_, err := kingpin.CommandLine.Parse(args)
	isTest, _ := strconv.ParseBool(*test)
	if err != nil {
		fmt.Fprintf(out, "Command line could not be parsed, error: %v", err)
		return
	}

	d, _ := ioutil.ReadFile(*data)
	puppies, uerr := unmarshalPuppies(d)
	if uerr != nil {
		return
	}

	switch *db {
	case "map":
		mapStore, err := storeInMap(puppies)
		if err != nil {
			fmt.Fprintf(out, "Could not setup map storage:%v", err)
		}
		fmt.Fprintf(out, "Map store of puppies:%v", mapStore)
		s = &mapStore
	case "sync":
		syncStore, err := storeInSync(puppies)
		if err != nil {
			fmt.Fprintf(out, "Could not setup sync storage:%v", err)
		}
		fmt.Fprintf(out, "Sync store of puppies:%v", syncStore)
		s = syncStore
	default:
		fmt.Fprintf(out, "Invalid storage")
	}

	if isTest {
		c <- 1
	}
	r := chi.NewRouter()

	rs := puppy.RestStorer{Db: s}
	r.Get(rs.GetPuppyRoute(), rs.GetPuppy)
	r.Post(rs.PostPuppyRoute(), rs.CreatePuppy)
	r.Put(rs.PutPuppyRoute(), rs.UpdatePuppy)
	r.Delete(rs.DeletePuppyRoute(), rs.DeletePuppy)

	portValue := fmt.Sprintf(":%v", *port)
	s := http.Server{Addr: portValue, Handler: r}
	serr := s.ListenAndServe()
	if serr != http.ErrServerClosed {
		fmt.Fprintf(out, "Could not start puppy server: %v", serr)
	}
}

func unmarshalPuppies(d []byte) ([]puppy.Puppy, error) {
	puppies := []puppy.Puppy{}
	if err := json.Unmarshal(d, &puppies); err != nil {
		fmt.Fprintf(out, "Could not unmarshall puppies, error: %v", err)
		return nil, err
	}
	return puppies, nil
}

func storeInMap(puppies []puppy.Puppy) (store.MapStore, error) {
	mapStore := store.NewMapStore()

	for _, p := range puppies {
		err := mapStore.CreatePuppy(p)
		if err != nil {
			fmt.Fprintf(out, "Could not create puppy, error: %v", err)
			return store.MapStore{}, err
		}
	}
	return *mapStore, nil
}

func storeInSync(puppies []puppy.Puppy) (*store.SyncStore, error) {
	syncStore := store.NewSyncStore()

	for _, p := range puppies {
		err := syncStore.CreatePuppy(p)
		if err != nil {
			fmt.Fprintf(out, "Could not create puppy, error: %v", err)
			return &store.SyncStore{}, err
		}
	}
	return syncStore, nil
}
