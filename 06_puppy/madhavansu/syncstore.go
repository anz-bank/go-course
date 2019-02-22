package main

/**
* Storer implementation in below scenarios
* Dynamic Map ID: MapStore implementation of Storer backed by a map
* Static Map ID: SyncStore implementation of Storer backed by a sync.Map
 */
//Sync store
func (s *syncStore) createPuppy(in Puppy) uint {
	s.Lock()
	defer s.Unlock()
	s.Store(in.id, in)
	return in.id
}

func (s *syncStore) readPuppy(id uint) Puppy {
	pd, ok := s.Load(id)
	if !ok {
		return Puppy{}
	}
	p, _ := pd.(Puppy)
	return p
}

func (s *syncStore) updatePuppy(id uint, in Puppy) {
	s.Store(in.id, in)
}

func (s *syncStore) deletePuppy(id uint) {
	s.Delete(id)
}
