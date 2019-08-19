package main

// NewSyncStore returns a pointer to a new instance of the SyncStore struct which implements the Storer interface.
func NewSyncStore() Storer {
	return &SyncStore{
		uuid: 0,
	}
}

// Create increments the uuid and adds the provided Puppy struct to the store with this identifier.
func (store *SyncStore) Create(puppy Puppy) (int, error) {
	if puppy.Value < 0 {
		return -1, NewError(NegativeValue)
	}

	puppy.ID = store.uuid
	store.Store(puppy.ID, puppy)
	store.uuid++

	return puppy.ID, nil
}

// Read returns the puppy matching the provided uuid.
// An empty Puppy struct is returned if the identifier does not exist.
func (store *SyncStore) Read(id int) (Puppy, error) {
	if value, ok := store.Load(id); ok {
		return value.(Puppy), nil
	}

	return Puppy{}, NewError(IDNotFound)
}

// Update modifies the puppy matching the provided uuid in the store with the provided Puppy struct.
// It returns a bool whether a matching identifier was modified or not.
func (store *SyncStore) Update(id int, puppy Puppy) (bool, error) {
	if _, ok := store.Load(id); !ok {
		return false, NewError(IDNotFound)
	}
	if puppy.Value < 0 {
		return false, NewError(NegativeValue)
	}

	puppy.ID = id
	store.Store(id, puppy)
	return true, nil
}

// Destroy removes the puppy matching the provided uuid from the store.
// It returns a bool whether a matching identifier was deleted or not.
func (store *SyncStore) Destroy(id int) (bool, error) {
	if _, ok := store.Load(id); !ok {
		return false, NewError(IDNotFound)
	}

	store.Delete(id)
	return true, nil
}
