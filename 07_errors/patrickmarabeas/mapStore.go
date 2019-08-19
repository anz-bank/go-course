package main

// NewMapStore returns a pointer to a new instance of the MapStore struct which implements the Storer interface.
func NewMapStore() Storer {
	return &MapStore{
		uuid:  0,
		store: map[int]Puppy{},
	}
}

// Create increments the uuid and adds the provided Puppy struct to the store with this identifier.
func (store *MapStore) Create(puppy Puppy) int {
	puppy.ID = store.uuid
	store.store[puppy.ID] = puppy
	store.uuid++

	return puppy.ID
}

// Read returns the puppy matching the provided uuid.
// An empty Puppy struct is returned if the identifier does not exist.
func (store *MapStore) Read(id int) Puppy {
	if _, ok := store.store[id]; ok {
		return store.store[id]
	}

	return Puppy{}
}

// Update modifies the puppy matching the provided uuid in the store with the provided Puppy struct.
// It returns a bool whether a matching identifier was modified or not.
func (store *MapStore) Update(id int, puppy Puppy) bool {
	if _, ok := store.store[id]; !ok {
		return false
	}

	puppy.ID = id
	store.store[id] = puppy
	return true
}

// Destroy removes the puppy matching the provided uuid from the store.
// It returns a bool whether a matching identifier was deleted or not.
func (store *MapStore) Destroy(id int) bool {
	if _, ok := store.store[id]; !ok {
		return false
	}

	delete(store.store, id)
	return true
}
