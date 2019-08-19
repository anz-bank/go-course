package main

// NewMapStore returns a pointer to a new instance of the MapStore struct which implements the Storer interface.
func NewMapStore() Storer {
	return &MapStore{
		uuid:  0,
		store: map[int]Puppy{},
	}
}

// Create increments the uuid and adds the provided Puppy struct to the store with this identifier.
func (store *MapStore) Create(puppy Puppy) (int, error) {
	if puppy.Value < 0 {
		return -1, NewError(NegativeValue)
	}

	puppy.ID = store.uuid
	store.store[puppy.ID] = puppy
	store.uuid++

	return puppy.ID, nil
}

// Read returns the puppy matching the provided uuid.
// An empty Puppy struct is returned if the identifier does not exist.
func (store *MapStore) Read(id int) (Puppy, error) {
	if _, ok := store.store[id]; ok {
		return store.store[id], nil
	}

	return Puppy{}, NewError(IDNotFound)
}

// Update modifies the puppy matching the provided uuid in the store with the provided Puppy struct.
// It returns a bool whether a matching identifier was modified or not.
func (store *MapStore) Update(id int, puppy Puppy) (bool, error) {
	if _, ok := store.store[id]; !ok {
		return false, NewError(IDNotFound)
	}
	if puppy.Value < 0 {
		return false, NewError(NegativeValue)
	}

	puppy.ID = id
	store.store[id] = puppy
	return true, nil
}

// Destroy removes the puppy matching the provided uuid from the store.
// It returns a bool whether a matching identifier was deleted or not.
func (store *MapStore) Destroy(id int) (bool, error) {
	if _, ok := store.store[id]; !ok {
		return false, NewError(IDNotFound)
	}

	delete(store.store, id)
	return true, nil
}
