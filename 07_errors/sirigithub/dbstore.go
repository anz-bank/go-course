package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/syndtr/goleveldb/leveldb"
)

type dbStore struct {
	dbfilePath string
	currID     int
}

func NewdbStore(filePath string) *dbStore {
	var store = dbStore{}
	store.dbfilePath = filePath
	return &store
}
func (s *dbStore) CreatePuppy(puppy *Puppy) (int, error) {
	if puppy.Value < 0 {
		return -1, NewError(ErrInvalidValue, "Puppy value must be greater than 0")
	}
	s.currID++
	puppy.ID = s.currID
	db := dbConn(s.dbfilePath)
	defer db.Close()

	value := getPuppyData(puppy)
	err := db.Put([]byte(strconv.Itoa(puppy.ID)), value, nil)
	if err != nil {
		return -1, NewError(ErrDatabaseWrite, fmt.Sprintf("Error creating Puppy %d in the dbStore", puppy.ID))
	}
	return puppy.ID, nil
}

// ReadPuppy reads an existing puppy from the store
func (s *dbStore) ReadPuppy(id int) (*Puppy, error) {
	db := dbConn(s.dbfilePath)
	defer db.Close()

	data, err := db.Get([]byte(strconv.Itoa(id)), nil)
	if err != nil {
		return nil, NewError(ErrIDNotFound, fmt.Sprintf("puppy ID %d does not exist in the database", id))
	}

	var p Puppy
	e := json.Unmarshal(data, &p)
	if e != nil {
		return nil, NewError(ErrUrmarshallData, fmt.Sprintf("Cannot marshall puppy value"))
	}
	return &p, nil
}
func (s *dbStore) UpdatePuppy(p *Puppy) error {
	db := dbConn(s.dbfilePath)
	defer db.Close()

	_, err := db.Get([]byte(strconv.Itoa(p.ID)), nil)
	if err != nil {
		println("error getting value from db")
		return NewError(ErrIDNotFound, fmt.Sprintf("puppy ID %d to update does not exist in the databse", p.ID))
	}
	value, err := json.Marshal(p)
	if err != nil {
		return NewError(ErrMarshallData, fmt.Sprintf("Cannot marshall puppy %d", p.ID))
	}
	e := db.Put([]byte(strconv.Itoa(p.ID)), value, nil)
	if e != nil {
		return NewError(ErrDatabaseWrite, fmt.Sprintf("Error writing Puppy %d to the databse", p.ID))
	}
	return nil
}

func (s *dbStore) DeletePuppy(id int) error {
	db := dbConn(s.dbfilePath)
	defer db.Close()
	err := db.Delete([]byte(strconv.Itoa(id)), nil)
	if err != nil {
		return NewError(ErrDatabseDelete, fmt.Sprintf("Error deleting Puppy %d from the databse", id))
	}
	return nil
}

func dbConn(dbfilePath string) (db *leveldb.DB) {
	db, err := leveldb.OpenFile("/tmp/foo.db", nil)
	if err != nil {
		panic(err)
	}
	return db
}

func getPuppyData(puppy *Puppy) []byte {
	value, err := json.Marshal(puppy)
	if err != nil {
		panic(err)
	}
	return value
}
