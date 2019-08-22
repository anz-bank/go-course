package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/suite"
)

type storerSuite struct {
	suite.Suite
	store      Storer
	storerType func() Storer
}

func TestStorers(t *testing.T) {
	suite.Run(t, &storerSuite{
		storerType: func() Storer { return InitMapStore() },
	})
	suite.Run(t, &storerSuite{
		storerType: func() Storer { return InitSyncStore() },
	})
}

func (s *storerSuite) SetupTest() {
	s.store = s.storerType()

	corgi := &Puppy{1000, "corgi", "orange and white", "$$$$$$$"}
	_ = s.store.CreatePuppy(corgi)
	spaniel := &Puppy{2000, "spaniel", "brown", "$$$$$$"}
	_ = s.store.CreatePuppy(spaniel)
	cujo := &Puppy{2500, "cujo", "tan", "$"}
	_ = s.store.CreatePuppy(cujo)
	terrier := &Puppy{3000, "terrier", "black", "$$$"}
	_ = s.store.CreatePuppy(terrier)
	bulldog := &Puppy{4000, "bulldog", "white", "$$"}
	_ = s.store.CreatePuppy(bulldog)
}

func (s *storerSuite) TestMapStore_CreatePuppy() {

	corgi := &Puppy{10000, "new corgi", "even brighter orange and white", "$$$$$$$"}

	tests := []struct {
		name    string
		puppy   *Puppy
		wantErr bool
	}{
		{
			name:    "New corgi",
			puppy:   corgi,
			wantErr: false,
		},
		{
			name:    "Existing corgi",
			puppy:   corgi,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		s.T().Run(tt.name, func(t *testing.T) {
			if err := s.store.CreatePuppy(tt.puppy); (err != nil) != tt.wantErr {
				t.Errorf("CreatePuppy() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func (s *storerSuite) TestMapStore_ReadPuppy() {

	corgi := &Puppy{1000, "corgi", "orange and white", "$$$$$$$"}
	_ = s.store.CreatePuppy(corgi)

	tests := []struct {
		name    string
		id      uint32
		wantErr bool
		want    *Puppy
	}{
		{
			name:    "lookup existing puppy",
			id:      1000,
			want:    corgi,
			wantErr: false,
		},
		{
			name:    "lookup non-existing puppy",
			id:      1001,
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		s.T().Run(tt.name, func(t *testing.T) {
			got, err := s.store.ReadPuppy(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadPuppy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("ReadPuppy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func (s *storerSuite) TestMapStore_DeletePuppy() {
	tests := []struct {
		name    string
		id      uint32
		want    bool
		wantErr bool
	}{
		{
			name:    "delete existing corgi",
			id:      1000,
			want:    true,
			wantErr: false,
		},
		{
			name:    "delete non-existent corgi",
			id:      1000,
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		s.T().Run(tt.name, func(t *testing.T) {
			got, err := s.store.DeletePuppy(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeletePuppy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != got {
				t.Errorf("DeletePuppy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func (s *storerSuite) TestMapStore_UpdatePuppy() {
	albinoCorgi := &Puppy{1000, "corgi", "white", "$$$$$$$"}
	corruptCorgi := &Puppy{1009, "corgi", "white", "$$$$$$$"}
	// _ = s.store.CreatePuppy(albinoCorgi)

	tests := []struct {
		name    string
		id      uint32
		wantErr bool
		puppy   *Puppy
	}{
		{
			name:    "update existing puppy",
			id:      1000,
			puppy:   albinoCorgi,
			wantErr: false,
		},
		{
			name:    "update non-existing puppy",
			id:      1001,
			puppy:   albinoCorgi,
			wantErr: true,
		},
		{
			name:    "update with an empty puppy",
			id:      1000,
			puppy:   nil,
			wantErr: true,
		},
		{
			name:    "update with a corrupt puppy",
			id:      1000,
			puppy:   corruptCorgi,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		s.T().Run(tt.name, func(t *testing.T) {
			if err := s.store.UpdatePuppy(tt.id, tt.puppy); (err != nil) != tt.wantErr {
				t.Errorf("UpdatePuppy() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
