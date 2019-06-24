package store

import (
	"encoding/json"
	"fmt"
	"strconv"
	"sync"

	puppy "github.com/anz-bank/go-course/08_project/n0npax/pkg/puppy"

	"github.com/syndtr/goleveldb/leveldb"
)

var levelDBPath = "/tmp/leveldb"

// LevelDBStore provides sync for leveldb
type LevelDBStore struct {
	ldb *leveldb.DB
	sync.Mutex
	nextID int
}

// NewLevelDBStorer creates new storer for leveldb
func NewLevelDBStorer() *LevelDBStore {
	db, err := leveldb.OpenFile(levelDBPath, nil)
	dbErrorPanic(err)
	return &LevelDBStore{nextID: 0, ldb: db}
}

// CreatePuppy creates puppy
func (l *LevelDBStore) CreatePuppy(p *puppy.Puppy) (int, error) {
	l.Lock()
	defer l.Unlock()
	id, err := l.putPuppy(l.nextID, p)
	l.nextID++
	return id, err
}

// ReadPuppy reads puppy from backend
func (l *LevelDBStore) ReadPuppy(id int) (*puppy.Puppy, error) {
	byteID := []byte(strconv.Itoa(id))
	if puppyData, err := l.ldb.Get(byteID, nil); err == nil {
		var p puppy.Puppy
		err := json.Unmarshal(puppyData, &p)
		if err != nil {
			return nil, puppy.Errorf(puppy.ErrInternalErrorCode, "Internal error. Could not cast stored data to puppy object")
		}
		return &p, nil
	}
	return nil, puppy.Errorf(puppy.ErrNotFoundCode, fmt.Sprintf("Puppy with ID (%v) not found", id))
}

// UpdatePuppy updates puppy
func (l *LevelDBStore) UpdatePuppy(id int, p *puppy.Puppy) error {
	if id != p.ID {
		return puppy.Errorf(puppy.ErrInvalidInputCode, "ID is corrupted. Please ensure object ID matched provided ID")
	}
	l.Lock()
	defer l.Unlock()
	if _, err := l.ReadPuppy(id); err != nil {
		return puppy.Errorf(puppy.ErrNotFoundCode, fmt.Sprintf("Puppy with ID (%v) not found", id))
	}
	_, err := l.putPuppy(id, p)
	return err
}

// DeletePuppy deletes puppy
func (l *LevelDBStore) DeletePuppy(id int) (bool, error) {
	l.Lock()
	defer l.Unlock()
	if _, err := l.ReadPuppy(id); err != nil {
		return false, puppy.Errorf(puppy.ErrNotFoundCode, fmt.Sprintf("Puppy with ID (%v) not found", id))
	}
	byteID := []byte(strconv.Itoa(id))
	err := l.ldb.Delete(byteID, nil)
	dbErrorPanic(err)
	return true, nil
}

// putPuppy stores puppy in backend
func (l *LevelDBStore) putPuppy(id int, p *puppy.Puppy) (int, error) {
	if p.Value < 0 {
		return -1, puppy.Errorf(puppy.ErrInvalidInputCode, "Puppy value have to be positive number")
	}
	puppyByte, _ := json.Marshal(p)
	byteID := []byte(strconv.Itoa(id))
	err := l.ldb.Put(byteID, puppyByte, nil)
	dbErrorPanic(err)
	return id, nil
}

// dbErrorPanic causes panic in error is not nil
func dbErrorPanic(err error) {
	if err != nil {
		panic(err)
	}
}
