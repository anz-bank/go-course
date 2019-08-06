package main

// NewSyncStore creates new storer for SyncMap
func NewSyncStore() *SyncStore {
	return &SyncStore{}
}

// CreatePuppy creates puppy
func (s *SyncStore) CreatePuppy(puppy *Puppy) (int, error) {
	if puppy.Value < 0 {
		return -1, Errorf(ErrCodeInvalidInput, "puppy value have to be positive number")
	}
	s.Lock()
	defer s.Unlock()
	puppy.ID = s.nextID
	s.nextID++
	s.Store(puppy.ID, *puppy)
	return puppy.ID, nil
}

// ReadPuppy reads puppy from backend
func (s *SyncStore) ReadPuppy(id int) (*Puppy, error) {
	if puppyData, ok := s.Load(id); ok {
		puppy := puppyData.(Puppy)
		return &puppy, nil
	}
	return nil, Errorf(ErrCodeNotFound, "puppy with ID (%v) not found", id)
}

// UpdatePuppy updates puppy
func (s *SyncStore) UpdatePuppy(id int, puppy *Puppy) error {
	if puppy.Value < 0 {
		return Errorf(ErrCodeInvalidInput, "puppy value have to be positive number")
	}
	if id != puppy.ID {
		return Errorf(ErrCodeInvalidInput, "ID is corrupted. Please ensure object ID matched provided ID")
	}
	s.Lock()
	defer s.Unlock()
	if _, ok := s.Load(id); !ok {
		return Errorf(ErrCodeNotFound, "puppy with ID (%v) not found", id)
	}
	s.Store(id, *puppy)
	return nil
}

// DeletePuppy deletes puppy
func (s *SyncStore) DeletePuppy(id int) (bool, error) {
	s.Lock()
	defer s.Unlock()
	if _, ok := s.Load(id); !ok {
		return false, Errorf(ErrCodeNotFound, "puppy with ID (%v) not found", id)
	}
	s.Delete(id)
	return true, nil
}
