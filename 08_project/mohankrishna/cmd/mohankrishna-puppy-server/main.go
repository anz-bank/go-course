package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	types "github.com/anz-bank/go-training/08_project/mohankrishna/pkg/mohankrishna-puppy"
	store "github.com/anz-bank/go-training/08_project/mohankrishna/pkg/mohankrishna-puppy/store"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var out io.Writer = os.Stdout

func parseArgs(args []string) {
	filePath := kingpin.Flag("data", "path to json data file").Short('d').String()
	_, err := kingpin.CommandLine.Parse(args)
	if err != nil {
		fmt.Fprint(out, err)
		return
	}
	b, err := ioutil.ReadFile(*filePath)
	if err != nil {
		fmt.Fprint(out, err)
		return
	}
	var puppies []types.Puppy
	err = json.Unmarshal(b, &puppies)
	if err != nil {
		fmt.Fprint(out, err)
		return
	}
	levelDBStore := store.NewLevelDBStore("level_store")
	for pos := range puppies {
		err = levelDBStore.CreatePuppy(&puppies[pos])
		if err != nil {
			fmt.Fprint(out, err)
			return
		}
	}
	pups, err := levelDBStore.GetAll()
	if err == nil {
		for _, pup := range pups {
			fmt.Fprint(out, *pup)
		}
	}
	levelDBStore.CloseDB()
}

func main() {
	parseArgs(os.Args[1:])
}
