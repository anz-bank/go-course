package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/alecthomas/kingpin"
	types "github.com/anz-bank/go-training/08_project/mohankrishna/pkg/mohankrishna-puppy"
	store "github.com/anz-bank/go-training/08_project/mohankrishna/pkg/mohankrishna-puppy/store"
)

var (
	dataFile     *string = kingpin.Flag("data", "path to json data file").Short('d').String()
	levelDBStore *store.LevelDBStore
)

func main() {
	kingpin.Parse()
	b, err := ioutil.ReadFile(*dataFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	var puppies []types.Puppy
	err = json.Unmarshal(b, &puppies)
	if err != nil {
		fmt.Println(err)
		return
	}
	levelDBStore = store.NewLevelDBStore()
	for pos := range puppies {
		err = levelDBStore.CreatePuppy(&puppies[pos])
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	pups, err := levelDBStore.GetAll()
	if err == nil {
		for _, pup := range pups {
			fmt.Println(*pup)
		}
	} else {
		fmt.Println(err)
		return
	}
	levelDBStore.CloseDB()
}
