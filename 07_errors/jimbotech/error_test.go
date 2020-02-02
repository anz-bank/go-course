package main

func (s *storesSuite) TestReadFailure() {
	pup2, err := s.store.ReadPuppy(1)
	s.Require().Nil(pup2)
	s.Require().Equal(ErrIDNotFound, err)
	s.Require().NotEmpty(err.Error())
}

func (s *storesSuite) TestNegReadFailure() {
	pup2, err := s.store.ReadPuppy(-1)
	s.Require().Nil(pup2)
	s.Require().Equal(ErrValueBelowZero, err)
}

func (s *storesSuite) TestUpdateError() {
	create(s)
	pup2 := Puppy{Breed: "kelpie", Colour: "black", Value: "indispensable"}
	err := s.store.UpdatePuppy(1, &pup2)
	s.Require().Equal(ErrIDNotFound, err)
}

func (s *storesSuite) TestNegUpdateError() {
	create(s)
	pup2 := Puppy{Breed: "kelpie", Colour: "black", Value: "indispensable"}
	err := s.store.UpdatePuppy(-1, &pup2)
	s.Require().Equal(ErrValueBelowZero, err)
}

func (s *storesSuite) TestNegDeleteFailure() {
	err := s.store.DeletePuppy(-1)
	s.Require().Equal(ErrValueBelowZero, err)
}

func (s *storesSuite) TestDeleteFailure() {
	err := s.store.DeletePuppy(1)
	s.Require().Equal(ErrIDNotFound, err)
}

func (s *storesSuite) TestError() {
	var err = &Error{}
	err = nil
	res := err.Error()
	s.Assert().Equal("<nil>", res)
}
