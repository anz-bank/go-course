package puppystorer

import "fmt"

// type MapStore map[int]Puppy
// This will serve as our in-memory DB
type MapStore struct {
	store  map[int]Puppy
	nextID int
}

func NewMapStore() *MapStore {
	return &MapStore{store: make(map[int]Puppy), nextID: 0}
}

// IncrementCounter increases the ID counter everytime a new Puppy is created to prevent overwrite issues
// in DeletePuppy()
func (ms *MapStore) incrementCounter() {
	ms.nextID++
}

// CreatePuppy lets you create a new unique puppy in MapStore
func (ms *MapStore) CreatePuppy(puppy *Puppy) int {
	ms.incrementCounter()
	puppy.ID = ms.nextID
	ms.store[puppy.ID] = *puppy // store the puppy object at this key in the map store (like a row in a table)
	return puppy.ID             // return u a handle to the correct row in the table that is MapStore
}

// ReadPuppy lets you GET a puppy from MapStore if it exists. Else it will return an error
func (ms *MapStore) ReadPuppy(id int) (*Puppy, error) {
	// if the key exists, then puppy is assigned the value stored under the key
	// referred to: https://blog.golang.org/go-maps-in-action
	if puppy, ok := ms.store[id]; ok {
		return &puppy, nil
	}
	// else return nil pointer to puppy and an error
	return nil, fmt.Errorf("puppy with ID: %d does not exist", id)
}

// UpdatePuppy lets you update a "row" in MapStore
func (ms *MapStore) UpdatePuppy(id int, puppy *Puppy) error {
	if _, ok := ms.store[id]; !ok {
		return fmt.Errorf("puppy with ID: %d does not exist", id)
	}
	puppy.ID = id // ignore id in puppy struct and use id passed as argument as id is created in storer
	ms.store[id] = *puppy
	return nil
}

// DeletePuppy lets you delete a specific puppy in MapStore
func (ms *MapStore) DeletePuppy(id int) error {
	if _, ok := ms.store[id]; !ok {
		return fmt.Errorf("puppy with ID: %d does not exist", id)
	}
	delete(ms.store, id)
	return nil
}
