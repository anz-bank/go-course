package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/anz-bank/go-course/10_rest/mohitnag/pkg/puppy"
	"github.com/anz-bank/go-course/10_rest/mohitnag/pkg/puppy/store"
	"github.com/anz-bank/go-course/10_rest/mohitnag/pkg/rest"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app   = kingpin.New("puppyStore", "Puppy Store")
	args  = os.Args[1:]
	srvCh chan *http.Server
)

func main() {
	var storeType string
	fileName := app.Flag("data", "file path").Short('d').ExistingFile()
	port := app.Flag("port", "port number").Short('p').Default("8686").String()
	app.Flag("store", "store type").Short('s').Default("sync").EnumVar(&storeType, "map", "sync")
	kingpin.MustParse(app.Parse(args))

	s := createStore(storeType)
	err := initialisePuppyStoreWithFile(s, *fileName)
	if err != nil {
		panic(err)
	}
	handler := rest.NewRestHandler(s)
	srv := &http.Server{
		Addr:    ":" + (*port),
		Handler: handler,
	}
	if srvCh != nil {
		srvCh <- srv
	}
	log.Panic(srv.ListenAndServe())
}

func initialisePuppyStoreWithFile(store store.Storer, fileName string) error {
	puppies := []puppy.Puppy{}
	puppiesBytes := readFile(fileName)
	if err := json.Unmarshal(puppiesBytes, &puppies); err != nil {
		return fmt.Errorf("could not unmarshal json: %v", err)
	}
	for _, puppy := range puppies {
		if err := store.CreatePuppy(puppy); err != nil {
			return err
		}
	}
	return nil
}

func readFile(filename string) []byte {
	buff, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return buff
}

func createStore(storeType string) store.Storer {
	switch storeType {
	case "map":
		return store.NewMapStore()
	default:
		return store.NewSyncStore()
	}
}
