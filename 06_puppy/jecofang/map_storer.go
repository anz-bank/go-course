package main

import "fmt"

type MapStore struct {
	m map[uint64]*Puppy
}

func NewMapStore() *MapStore {
	return &MapStore{make(map[uint64]*Puppy)}
}

func (storer *MapStore) Create(puppy *Puppy) error {
	if err := validateID(puppy.ID); err != nil {
		return err
	}
	if _, exists := storer.m[puppy.ID]; exists {
		return &StorerError{Conflict, fmt.Sprintf("The Puppy with ID `%d` already exists.", puppy.ID)}
	}
	storer.m[puppy.ID] = puppy
	return nil
}

func (storer *MapStore) Read(id uint64) (*Puppy, error) {
	if err := validateID(id); err != nil {
		return nil, err
	}
	if puppy, exists := storer.m[id]; exists {
		return puppy, nil
	}
	return nil, &StorerError{NotFound, fmt.Sprintf("The puppy with ID `%d` does not exist.", id)}
}

func (storer *MapStore) Update(id uint64, puppy *Puppy) error {
	if err := validateID(id); err != nil {
		return err
	}
	if err := validateID(puppy.ID); err != nil {
		return &StorerError{Invalid, fmt.Sprintf("The input %v has invalid ID `%d`.", *puppy, puppy.ID)}
	}
	if id != puppy.ID {
		return &StorerError{Invalid,
			fmt.Sprintf("The ID mismatch; The given id is `%d` but the puppy.ID is `%d`.", id, puppy.ID)}
	}
	if _, exists := storer.m[id]; !exists {
		return &StorerError{NotFound, fmt.Sprintf("The puppy with ID `%d` does not exist.", id)}
	}
	storer.m[id] = puppy
	return nil
}

func (storer *MapStore) Delete(id uint64) error {
	if err := validateID(id); err != nil {
		return err
	}
	if _, exists := storer.m[id]; exists {
		delete(storer.m, id)
		return nil
	}
	return &StorerError{NotFound, fmt.Sprintf("The puppy with ID `%d` does not exist.", id)}
}
