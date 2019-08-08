package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	puppy "github.com/anz-bank/go-course/09_json/runnerdave/pkg/puppy"
	store "github.com/anz-bank/go-course/09_json/runnerdave/pkg/puppy/store"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	args           = os.Args[1:]
	data           = kingpin.Flag("data", "data file").Short('d').Default("data.json").ExistingFile()
	out  io.Writer = os.Stdout
)

func main() {
	_, err := kingpin.CommandLine.Parse(args)
	if err != nil {
		fmt.Fprintf(out, "Command line could not be parsed, error: %v", err)
		return
	}

	d, _ := ioutil.ReadFile(*data)
	puppies, uerr := unmarshalPuppies(d)
	if uerr != nil {
		return
	}

	mapStore := storeInMap(puppies)
	fmt.Fprintf(out, "Map store of puppies:%v", mapStore)

	syncStore := storeInSync(puppies)
	fmt.Fprintf(out, "Sync store of puppies:%v", syncStore)
}

func unmarshalPuppies(d []byte) ([]puppy.Puppy, error) {
	puppies := []puppy.Puppy{}
	if err := json.Unmarshal(d, &puppies); err != nil {
		fmt.Fprintf(out, "Could not unmarshall puppies, error: %v", err)
		return nil, err
	}
	return puppies, nil
}

func storeInMap(puppies []puppy.Puppy) store.MapStore {
	mapStore := store.NewMapStore()

	for _, p := range puppies {
		err := mapStore.CreatePuppy(p)
		if err != nil {
			fmt.Fprintf(out, "Could not create puppy, error: %v", err)
			return store.MapStore{}
		}
	}
	return *mapStore
}

func storeInSync(puppies []puppy.Puppy) *store.SyncStore {
	syncStore := store.NewSyncStore()

	for _, p := range puppies {
		err := syncStore.CreatePuppy(p)
		if err != nil {
			fmt.Fprintf(out, "Could not create puppy, error: %v", err)
		}
	}
	return syncStore
}
