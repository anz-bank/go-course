package main

import (
	"encoding/binary"
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

type LevelDbStore struct {
	conn string
}

func NewLevelDbStore(dbFilePath string) *LevelDbStore {
	return &LevelDbStore{dbFilePath}
}

func (storer *LevelDbStore) Create(puppy *Puppy) error {
	if err := validateID(puppy.ID); err != nil {
		return err
	}

	db, _ := leveldb.OpenFile(storer.conn, nil)
	defer db.Close()

	key := storer.formatKey(puppy.ID)
	if exists, _ := db.Has(key, nil); exists {
		return &StorerError{Conflict, fmt.Sprintf("The Puppy with ID `%d` already exists.", puppy.ID)}
	}

	err := db.Put(key, puppy.ToJSON(), nil)
	return err
}

func (storer *LevelDbStore) Read(id uint64) (*Puppy, error) {
	if err := validateID(id); err != nil {
		return nil, err
	}

	db, _ := leveldb.OpenFile(storer.conn, nil)
	defer db.Close()

	key := storer.formatKey(id)
	data, err := db.Get(key, nil)
	if err != nil {
		return nil, &StorerError{NotFound, fmt.Sprintf("The puppy with ID `%d` does not exist.", id)}
	}

	var puppy Puppy
	err = puppy.ParseJSON(data)
	return &puppy, err
}

func (storer *LevelDbStore) Update(id uint64, puppy *Puppy) error {
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

	db, _ := leveldb.OpenFile(storer.conn, nil)
	defer db.Close()

	key := storer.formatKey(id)
	if exists, _ := db.Has(key, nil); !exists {
		return &StorerError{NotFound, fmt.Sprintf("The puppy with ID `%d` does not exist.", id)}
	}

	err := db.Put(key, puppy.ToJSON(), nil)
	return err
}

func (storer *LevelDbStore) Delete(id uint64) error {
	if err := validateID(id); err != nil {
		return err
	}

	db, _ := leveldb.OpenFile(storer.conn, nil)
	defer db.Close()

	key := storer.formatKey(id)
	if exists, _ := db.Has(key, nil); !exists {
		return &StorerError{NotFound, fmt.Sprintf("The puppy with ID `%d` does not exist.", id)}
	}

	err := db.Delete(key, nil)
	return err
}

func (storer *LevelDbStore) formatKey(id uint64) []byte {
	buf := make([]byte, binary.MaxVarintLen64)
	n := binary.PutUvarint(buf, id)
	return buf[:n]
}
