package main

/**
* Storer implementation in below scenarios
* Dynamic Map ID: MapStore implementation of Storer backed by a map
* Static Map ID: SyncStore implementation of Storer backed by a sync.Map
 */
// Map Store
func (m *mapStore) createPuppy(in Puppy) uint {
	m.mapID++
	in.id = m.mapID
	m.ms[in.id] = in
	return m.mapID
}

func (m *mapStore) readPuppy(id uint) Puppy {
	p, ok := m.ms[id]
	if !ok {
		return Puppy{}
	}
	return p
}

func (m *mapStore) updatePuppy(id uint, in Puppy) {
	m.ms[id] = in
}

func (m *mapStore) deletePuppy(id uint) {
	delete(m.ms, id)
}
