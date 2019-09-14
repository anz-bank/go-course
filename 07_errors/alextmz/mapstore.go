package main

type MapStore map[PuppyID]*Puppy

func NewmapStore() *MapStore {
	a := MapStore{}
	return &a
}

func (m *MapStore) CreatePuppy(p *Puppy) error {
	if _, ok := (*m)[p.ID]; !ok {
		if p.Value < 0 {
			return NewError(ErrValueLessThanZero)
		}
		(*m)[p.ID] = p
		return nil
	}
	return NewError(ErrIDBeingCreatedAlreadyExists)
}

func (m *MapStore) ReadPuppy(id PuppyID) (*Puppy, error) {
	if v, ok := (*m)[id]; ok {
		return v, nil
	}
	return nil, NewError(ErrIDBeingReadDoesNotExist)
}

func (m *MapStore) UpdatePuppy(id PuppyID, p *Puppy) error {
	if _, ok := (*m)[id]; ok {
		if p.Value < 0 {
			return NewError(ErrValueLessThanZero)
		}
		(*m)[id] = p
		return nil
	}
	return NewError(ErrIDBeingUpdatedDoesNotExist)

}

func (m *MapStore) DeletePuppy(id PuppyID) error {
	if _, ok := (*m)[id]; ok {
		delete(*m, id)
		return nil
	}
	return NewError(ErrIDBeingDeletedDoesNotExist)
}
