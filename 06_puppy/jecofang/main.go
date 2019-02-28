package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

var out io.Writer = os.Stdout

func main() {
	p1 := &Puppy{1, "Beagle", "Brown", 132.23}
	p2 := &Puppy{1, "Boxer", "Black", 2000.00}

	var ms Storer = NewMapStore()
	fmt.Fprintln(out, "Map Store")

	if err := ms.Create(p1); err == nil {
		fmt.Fprintf(out, "Create %v succeed\n", p1)
	}

	if r, err := ms.Read(1); err == nil {
		fmt.Fprintf(out, "Read %v with ID %d succeed\n", r, 1)
	}

	if err := ms.Update(1, p2); err == nil {
		fmt.Fprintf(out, "Update %v succeed\n", p2)
	}

	if r, err := ms.Read(1); err == nil {
		fmt.Fprintf(out, "Read %v with ID %d succeed\n", r, 1)
	}

	if err := ms.Delete(1); err == nil {
		fmt.Fprintf(out, "Delete Puppy with ID %d succeed\n", 1)
	}

	var sm Storer = NewSyncStore()
	fmt.Fprintln(out, "SyncMap Store")

	if err := sm.Create(p1); err == nil {
		fmt.Fprintf(out, "Create %v succeed\n", p1)
	}

	if r, err := sm.Read(1); err == nil {
		fmt.Fprintf(out, "Read %v with ID %d succeed\n", r, 1)
	}

	if err := sm.Update(1, p2); err == nil {
		fmt.Fprintf(out, "Update %v succeed\n", p2)
	}

	if r, err := sm.Read(1); err == nil {
		fmt.Fprintf(out, "Read %v with ID %d succeed\n", r, 1)
	}

	if err := sm.Delete(1); err == nil {
		fmt.Fprintf(out, "Delete Puppy with ID %d succeed\n", 1)
	}

	dbPath := filepath.Join(os.TempDir(), "leveldb-puppy")

	var sl Storer = NewLevelDbStore(dbPath)
	fmt.Fprintln(out, "levelDB Store")

	if err := sl.Create(p1); err == nil {
		fmt.Fprintf(out, "Create %v succeed\n", p1)
	}

	if r, err := sl.Read(1); err == nil {
		fmt.Fprintf(out, "Read %v with ID %d succeed\n", r, 1)
	}

	if err := sl.Update(1, p2); err == nil {
		fmt.Fprintf(out, "Update %v succeed\n", p2)
	}

	if r, err := sl.Read(1); err == nil {
		fmt.Fprintf(out, "Read %v with ID %d succeed\n", r, 1)
	}

	if err := sl.Delete(1); err == nil {
		fmt.Fprintf(out, "Delete Puppy with ID %d succeed\n", 1)
	}

}
