package store

import (
	"fmt"

	"github.com/anz-bank/go-course/08_project/nickolee/pkg/puppy"
)

// NewMapStore conveniently creates a new initialised syncstore
func NewMapStore() *MapStore {
	return &MapStore{store: make(map[int]puppy.Puppy), nextID: 0}
}

// MapStore will serve as our in-memory DB
type MapStore struct {
	store  map[int]puppy.Puppy
	nextID int
}

// IncrementCounter increases the ID counter everytime a new Puppy is created to prevent overwrite issues
// in DeletePuppy()
func (ms *MapStore) incrementCounter() {
	ms.nextID++
}

// CreatePuppy lets you create a new unique puppy in MapStore
func (ms *MapStore) CreatePuppy(pup *puppy.Puppy) (int, error) {
	// Check for negative value. If negative return custom error type
	if pup.Value < 0 {
		return 0, &puppy.Error{
			Message: "Sorry puppy value cannot be negative. The dog has to be worth something :)",
			Code:    puppy.ErrNegativePuppyID,
		}
	}

	// Else create new puppy (happy path)
	ms.incrementCounter()
	pup.ID = ms.nextID
	ms.store[pup.ID] = *pup // store the puppy object at this key in the map store (like a row in a table)
	return pup.ID, nil      // return u a handle to the correct row in the table that is MapStore
}

// ReadPuppy lets you GET a puppy from MapStore if it exists. Else it will return an error
func (ms *MapStore) ReadPuppy(id int) (*puppy.Puppy, error) {
	// if the key exists, then puppy is assigned the value stored under the key
	// referred to: https://blog.golang.org/go-maps-in-action
	if pup, ok := ms.store[id]; ok {
		return &pup, nil
	}

	// else return nil pointer to puppy and one of our custom errors
	return nil, &puppy.Error{
		Message: fmt.Sprintf("Sorry puppy with ID %d does not exist", id),
		Code:    puppy.ErrPuppyNotFound,
	}
}

// UpdatePuppy lets you update a "row" in MapStore
func (ms *MapStore) UpdatePuppy(id int, pup *puppy.Puppy) error {
	if _, ok := ms.store[id]; !ok {
		return &puppy.Error{
			Message: fmt.Sprintf("Sorry puppy with ID %d does not exist", id),
			Code:    puppy.ErrPuppyNotFound,
		}
	}

	pup.ID = id // ignore id in puppy struct and use id passed as argument as id is created in storer
	ms.store[id] = *pup
	return nil
}

// DeletePuppy lets you delete a specific puppy in MapStore
func (ms *MapStore) DeletePuppy(id int) error {
	if _, ok := ms.store[id]; !ok {
		return &puppy.Error{
			Message: fmt.Sprintf("Sorry puppy with ID %d does not exist", id),
			Code:    puppy.ErrPuppyNotFound,
		}
	}
	delete(ms.store, id)
	return nil
}
