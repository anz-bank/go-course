package main

import (
	"fmt"
	"io"
	"os"

	puppy "github.com/anz-bank/go-course/08_project/willshen8/pkg/puppy"
	store "github.com/anz-bank/go-course/08_project/willshen8/pkg/puppy/store"
)

var out io.Writer = os.Stdout

func main() {
	cutePuppy := puppy.Puppy{ID: 1, Breed: "Jack Russell", Color: "White and Brown", Value: "500"}
	naughtyPuppy := puppy.Puppy{ID: 1234, Breed: "Fox Terrier", Color: "Black", Value: "1300"}

	var mystore = store.NewMapStore()

	createdID, _ := mystore.CreatePuppy(&cutePuppy)
	fmt.Fprintln(out, "First Created Puppy ID:", createdID)

	readPuppy, _ := mystore.ReadPuppy(createdID)
	fmt.Fprintln(out, "ReadPuppy:", readPuppy)

	updateResult, _ := mystore.UpdatePuppy(createdID, &naughtyPuppy)
	fmt.Fprintln(out, "Update puppy result:", updateResult)

	deleteSuccess, _ := mystore.DeletePuppy(createdID)
	fmt.Fprintln(out, "Result of deleting puppy:", deleteSuccess)
}
