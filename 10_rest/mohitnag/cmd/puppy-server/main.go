package main

import (
	"encoding/json"
	"io/ioutil"
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
	srvCh = make(chan *http.Server)
)

func main() {
	var storeType string
	fileName := app.Flag("data", "file path").Short('d').ExistingFile()
	port := app.Flag("port", "port number").Short('p').Default("8686").String()
	app.Flag("store", "store type").Short('s').Default("sync").EnumVar(&storeType, "map", "sync")
	kingpin.MustParse(app.Parse(args))

	s, err := initialisePuppyStoreWithFile(storeType, *fileName)
	if err != nil {
		panic(err)
	}
	handler := rest.NewRestHandler(s)
	srv := &http.Server{
		Addr:    ":" + (*port),
		Handler: handler,
	}
	srvCh <- srv
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}

func initialisePuppyStoreWithFile(storeType string, fileName string) (puppy.Storer, error) {
	store := createStore(storeType)
	puppies := []puppy.Puppy{}
	puppiesBytes := readFile(fileName)
	if err := json.Unmarshal(puppiesBytes, &puppies); err != nil {
		panic(err)
	}
	for _, puppy := range puppies {
		if err := store.CreatePuppy(puppy); err != nil {
			return nil, err
		}
	}
	return store, nil
}

func readFile(filename string) []byte {
	buff, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return buff
}

func createStore(storeType string) puppy.Storer {
	switch storeType {
	case "map":
		return store.NewMapStore()
	default:
		return store.NewSyncStore()
	}
}
