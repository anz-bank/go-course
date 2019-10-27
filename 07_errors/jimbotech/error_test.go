package main

func (s *storesSuite) TestUpdateError() {
	_, pup := create(s)
	pup2 := Puppy{Breed: "kelpie", Colour: "black", Value: "indispensable"}
	err := s.store.UpdatePuppy(-1, &pup2)
	s.Require().IsType(&Error{}, err)
	actualErr, _ := err.(*Error)
	s.Require().Equal(ErrNegativeID, actualErr.Code)
	// now check by reading the updated value back and compare
	pup3, err2 := s.store.ReadPuppy(pup.ID)
	if s.Nil(err2, "Reading back updated value should work") {
		s.NotEqual(pup2, *pup3)
	}
}

func (s *storesSuite) TestReadFailure() {
	pup2, err := s.store.ReadPuppy(1)
	s.Require().Nil(pup2)
	s.Require().IsType(&Error{}, err)
	actualErr, _ := err.(*Error)
	s.Assert().Equal(ErrNotFound, actualErr.Code)
}

func (s *storesSuite) TestNegDeleteFailure() {
	err := s.store.DeletePuppy(1)
	s.Require().IsType(&Error{}, err)
	actualErr, _ := err.(*Error)
	s.Assert().Equal(ErrNotFound, actualErr.Code)
}

func (s *storesSuite) TestDeleteFailure() {
	err := s.store.DeletePuppy(-1)
	s.Require().IsType(&Error{}, err)
	actualErr, _ := err.(*Error)
	s.Assert().Equal(ErrNegativeID, actualErr.Code)
	s.Assert().NotEmpty(actualErr.Error())
}

func (s *storesSuite) TestError() {
	var err = &Error{}
	err = nil
	res := err.Error()
	s.Assert().Equal("<nil>", res)
}
