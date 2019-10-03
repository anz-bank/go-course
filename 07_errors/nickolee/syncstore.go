package puppystorer

import (
	"fmt"
	"sync"
)

// NewSyncStore() conveniently creates a new initialised syncstore
func NewSyncStore() *SyncStore {
	return &SyncStore{Map: sync.Map{}}
}

// SyncStore struct. To serve as alternative in-memory DB. It also implements Storer interface
type SyncStore struct {
	sync.Map
	nextID int
	sync.Mutex
}

// IncrementCounter increases the ID counter everytime a new Puppy is created to prevent overwrite issues
// in DeletePuppy()
func (m *SyncStore) incrementCounter() {
	m.nextID++
}

// CreatePuppy creates a puppy in sync store. Note we use a pointer receiver for this
func (m *SyncStore) CreatePuppy(puppy *Puppy) (int, error) {
	// Check for negative value. If negative return custom error type
	if puppy.Value < 0 {
		return 0, &Error{
			Message: "Sorry puppy value cannot be negative. The dog has to be worth something :)",
			Code:    ErrNegativePuppyID,
		}
	}

	// Add locking for thread safety before a write operation occurs
	m.Lock()
	defer m.Unlock()

	// Else create new puppy (happy path)
	m.incrementCounter()
	puppy.ID = m.nextID
	m.Store(puppy.ID, *puppy)
	return puppy.ID, nil
}

// ReadPuppy retrieves puppy from sync store
func (m *SyncStore) ReadPuppy(id int) (*Puppy, error) {
	if puppyData, exists := m.Load(id); exists {
		// This is not to do with calling method or accessing field, it's saying "cast to puppy"
		puppy := puppyData.(Puppy)
		return &puppy, nil
	}

	// else return nil pointer to puppy and one of our custom errors
	return nil, &Error{
		Message: fmt.Sprintf("Sorry puppy with ID %d does not exist", id),
		Code:    ErrPuppyNotFound,
	}
}

// UpdatePuppy updates a puppy in sync store
func (m *SyncStore) UpdatePuppy(id int, puppy *Puppy) error {
	// Add locking for thread safety before a write operation occurs
	m.Lock()
	defer m.Unlock()

	if _, exists := m.Load(id); !exists {
		return &Error{
			Message: fmt.Sprintf("Sorry puppy with ID %d does not exist", id),
			Code:    ErrPuppyNotFound,
		}
	}

	puppy.ID = id // ignore id in puppy struct and use id passed as argument as id is created in storer
	m.Store(id, *puppy)
	return nil
}

// DeletePuppy deletes a puppy in sync store
func (m *SyncStore) DeletePuppy(id int) error {
	// Add locking for thread safety before a write operation occurs
	m.Lock()
	defer m.Unlock()

	if _, exists := m.Load(id); !exists {
		return &Error{
			Message: fmt.Sprintf("Sorry puppy with ID %d does not exist", id),
			Code:    ErrPuppyNotFound,
		}
	}
	m.Delete(id)
	return nil
}
