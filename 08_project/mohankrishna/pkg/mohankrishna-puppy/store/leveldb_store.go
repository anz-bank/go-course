package store

import (
	"encoding/json"
	"fmt"

	types "github.com/anz-bank/go-training/08_project/mohankrishna/pkg/mohankrishna-puppy"
	"github.com/syndtr/goleveldb/leveldb"
)

type LevelDBStore struct {
	ldb *leveldb.DB
}

func NewLevelDBStore() *LevelDBStore {
	db, _ := leveldb.OpenFile("storage", nil)
	return &LevelDBStore{db}
}

func (ls *LevelDBStore) CloseDB() {
	ls.ldb.Close()
}

func (ls *LevelDBStore) CreatePuppy(puppy *types.Puppy) error {
	b, _ := json.Marshal(puppy.ID)
	ok, _ := ls.ldb.Has(b, nil)
	if ok {
		return &types.Error{Code: types.ErrDuplicate,
			Message: fmt.Sprintf("A puppy with ID: %d already exists", puppy.ID)}
	}
	bytes, _ := json.Marshal(puppy)
	err := ls.ldb.Put(b, bytes, nil)
	return err
}

func (ls *LevelDBStore) ReadPuppy(id uint32) (*types.Puppy, error) {
	b, _ := json.Marshal(id)
	data, err := ls.ldb.Get(b, nil)
	if err != nil {
		return nil, err
	}
	var puppy types.Puppy
	err = json.Unmarshal(data, &puppy)
	return &puppy, err
}

func (ls *LevelDBStore) UpdatePuppy(id uint32, puppy *types.Puppy) error {
	if id != puppy.ID {
		return &types.Error{Code: types.ErrInvalidInput,
			Message: fmt.Sprintf("The id:%d passed and the puppy's id:%d doesnt match", id, puppy.ID)}
	}
	b, _ := json.Marshal(id)
	bytes, _ := json.Marshal(puppy)
	err := ls.ldb.Put(b, bytes, nil)
	return err
}

func (ls *LevelDBStore) DeletePuppy(id uint32) error {
	b, _ := json.Marshal(id)
	ok, _ := ls.ldb.Has(b, nil)
	if !ok {
		return &types.Error{Code: types.ErrNotFound, Message: fmt.Sprintf("No puppy exists with id %d", id)}
	}
	return ls.ldb.Delete(b, nil)
}

func (ls *LevelDBStore) GetAll() ([]*types.Puppy, error) {
	var puppies []*types.Puppy
	iter := ls.ldb.NewIterator(nil, nil)
	for iter.Next() {
		var pup types.Puppy
		err := json.Unmarshal(iter.Value(), &pup)
		if err == nil {
			puppies = append(puppies, &pup)
		}
	}
	iter.Release()
	return puppies, iter.Error()
}
