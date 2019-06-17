package store

import (
	"sync"

	puppy "github.com/anz-bank/go-course/08_project/n0npax/pkg/puppy"

	"github.com/syndtr/goleveldb/leveldb"
)

//SyncStore sync.Map based type for storing puppies data
type SyncStore struct {
	sync.Map
	sync.Mutex
	total int
}

// MemStore map based type for storing puppies data
type MemStore map[int]puppy.Puppy

// LevelDBStore provides sync for leveldb
type LevelDBStore struct {
	ldb *leveldb.DB
	sync.Mutex
	total int
}
