package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/anz-bank/go-course/11_notify/kasunfdo/pkg/puppy"
	"github.com/anz-bank/go-course/11_notify/kasunfdo/pkg/puppy/store"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// Config contains puppyserver configurations.
type Config struct {
	dataFile   io.Reader
	port       int
	storeType  string
	lostSvcURL string
}

func main() {
	config, err := parseArgs(os.Args[1:])
	if err != nil {
		log.Fatalf("Failed to parse args.\n%v", err)
	}

	log.Fatal(initPuppyServer(config))
}

func parseArgs(args []string) (Config, error) {
	dataFile := kingpin.Flag("data", "Path to puppy data file").Short('d').Required().File()
	port := kingpin.Flag("port", "Puppy server port").Short('p').Default("8080").Int()
	storeType := kingpin.Flag("store", "Puppy store type").Short('s').Default("map").Enum("map", "sync")
	lostSvcURL := kingpin.Flag("lostsvc", "Lost puppy service endpoint").Short('l').Required().String()
	_, err := kingpin.CommandLine.Parse(args)

	return Config{*dataFile, *port, *storeType, *lostSvcURL}, err
}

func initPuppyServer(config Config) error {
	puppyStore, err := createStore(config.storeType)
	if err != nil {
		return err
	}
	err = createPuppies(config.dataFile, puppyStore)
	if err != nil {
		return err
	}

	apiHandler := puppy.NewAPIHandler(puppyStore, config.lostSvcURL)
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	apiHandler.WireRoutes(router)

	serverAddr := fmt.Sprintf(":%d", config.port)
	return http.ListenAndServe(serverAddr, router)
}

func createStore(storeType string) (puppy.Storer, error) {
	switch storeType {
	case "map":
		return store.NewMapStore(), nil
	case "sync":
		return store.NewSyncStore(), nil
	default:
		return nil, puppy.ErrorEf(puppy.ErrInvalid, nil, "given store type %v is invalid", storeType)
	}
}

func createPuppies(dataFile io.Reader, store puppy.Storer) error {
	data, err := ioutil.ReadAll(dataFile)
	if err != nil {
		return puppy.ErrorEf(puppy.ErrInternal, err, "failed to read puppy data file")
	}

	var puppies []puppy.Puppy
	if err = json.Unmarshal(data, &puppies); err != nil {
		return puppy.ErrorEf(puppy.ErrBadFormat, err, "failed to decode")
	}

	for _, p := range puppies {
		var id uint64
		if id, err = store.CreatePuppy(p); err != nil {
			return err
		}
		fmt.Printf("%v (id: %d) added to store\n", p, id)
	}

	return nil
}
