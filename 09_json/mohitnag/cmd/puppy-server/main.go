package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/anz-bank/go-course/09_json/mohitnag/pkg/puppy"
	"github.com/anz-bank/go-course/09_json/mohitnag/pkg/puppy/store"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	out  io.Writer = os.Stdout
	app            = kingpin.New("puppyStore", "Puppy Store")
	args           = os.Args[1:]
)

func main() {
	mapStore := store.MapStore{}
	syncStore := store.SyncStore{}
	fileName := app.Flag("data", "file path").Short('d').ExistingFile()
	kingpin.MustParse(app.Parse(args))
	if err := initialisePuppyStore(&mapStore, &syncStore, *fileName); err != nil {
		panic(err)
	}
	puppyMapStore, _ := mapStore.ReadPuppy(1)
	puppySyncStore, _ := syncStore.ReadPuppy(1)
	fmt.Fprintln(out, puppyMapStore)
	fmt.Fprintln(out, puppySyncStore)
}

func initialisePuppyStore(m *store.MapStore, s *store.SyncStore, fileName string) error {
	puppies := puppy.Puppies{}
	puppiesString := readJSON(fileName)
	if err := json.Unmarshal([]byte(puppiesString), &puppies); err != nil {
		panic(err)
	}
	for _, puppy := range puppies {
		if err := m.CreatePuppy(puppy); err != nil {
			return err
		}
		if err := s.CreatePuppy(puppy); err != nil {
			return err
		}
	}
	return nil
}

func readJSON(filename string) string {
	buff, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(buff)
}