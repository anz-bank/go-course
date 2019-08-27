package puppystorer

import (
	"fmt"
	"sync"
)

// SyncStore struct. To serve as alternative in-memory DB. It also implements Storer interface
type SyncStore struct {
	sync.Map
	nextID int
}

func NewSyncStore() *SyncStore {
	return &SyncStore{Map: sync.Map{}}
}

// IncrementCounter increases the ID counter everytime a new Puppy is created to prevent overwrite issues
// in DeletePuppy()
func (m *SyncStore) incrementCounter() {
	m.nextID++
}

// CreatePuppy creates a puppy in sync store. Note we use a pointer receiver for this
func (m *SyncStore) CreatePuppy(puppy *Puppy) int {
	m.incrementCounter()
	puppy.ID = m.nextID
	m.Store(puppy.ID, *puppy)
	return puppy.ID
}

// ReadPuppy retrieves puppy from sync store
func (m *SyncStore) ReadPuppy(id int) (*Puppy, error) {
	if puppyData, exists := m.Load(id); exists {
		// This is not to do with calling method or accessing field, it's saying "cast to puppy"
		puppy := puppyData.(Puppy)
		return &puppy, nil
	}
	return nil, fmt.Errorf("puppy with ID: %d does not exist", id)
}

// UpdatePuppy updates a puppy in sync store
func (m *SyncStore) UpdatePuppy(id int, puppy *Puppy) error {
	if _, exists := m.Load(id); !exists {
		return fmt.Errorf("puppy with ID: %d does not exist", id)
	}
	puppy.ID = id // ignore id in puppy struct and use id passed as argument as id is created in storer
	m.Store(id, *puppy)
	return nil
}

// DeletePuppy deletes a puppy in sync store
func (m *SyncStore) DeletePuppy(id int) error {
	if _, exists := m.Load(id); !exists {
		return fmt.Errorf("puppy with ID: %d does not exist", id)
	}
	m.Delete(id)
	return nil
}
