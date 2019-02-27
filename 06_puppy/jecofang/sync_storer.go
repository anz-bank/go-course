package main

import (
	"fmt"
	"sync"
)

type SyncStore struct {
	m sync.Map
}

func NewSyncStore() *SyncStore {
	return &SyncStore{}
}

func (storer *SyncStore) Create(puppy *Puppy) error {
	if err := validateID(puppy.ID); err != nil {
		return err
	}
	if _, exists := storer.m.Load(puppy.ID); exists {
		return &StorerError{Conflict, fmt.Sprintf("The Puppy with ID `%d` already exists.", puppy.ID)}
	}
	storer.m.Store(puppy.ID, puppy)
	return nil
}

func (storer *SyncStore) Read(id uint64) (*Puppy, error) {
	if err := validateID(id); err != nil {
		return nil, err
	}
	if v, exists := storer.m.Load(id); exists {
		if puppy, ok := v.(*Puppy); ok {
			return puppy, nil
		}
	}

	return nil, &StorerError{NotFound, fmt.Sprintf("The puppy with ID `%d` does not exist.", id)}
}

func (storer *SyncStore) Update(id uint64, puppy *Puppy) error {
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
	if _, exists := storer.m.Load(id); !exists {
		return &StorerError{NotFound, fmt.Sprintf("The puppy with ID `%d` does not exist.", id)}
	}
	storer.m.Store(id, puppy)
	return nil
}

func (storer *SyncStore) Delete(id uint64) error {
	if err := validateID(id); err != nil {
		return err
	}
	if _, exists := storer.m.Load(id); exists {
		storer.m.Delete(id)
		return nil
	}
	return &StorerError{NotFound, fmt.Sprintf("The puppy with ID `%d` does not exist.", id)}
}
