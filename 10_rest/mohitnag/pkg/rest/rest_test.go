package rest

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/anz-bank/go-course/10_rest/mohitnag/pkg/puppy"
	"github.com/anz-bank/go-course/10_rest/mohitnag/pkg/puppy/store"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type restSuite struct {
	suite.Suite
	handler    *Handler
	makeStorer func() store.Storer
}

func (s *restSuite) SetupTest() {
	s.handler = NewRestHandler(s.makeStorer()).(*Handler)
	err := initialisePuppyStore(s.handler.storer, "./../puppy/store/testdata/many-puppies.json")
	if err != nil {
		s.FailNow("Error in test setup")
	}
}

func TestStorers(t *testing.T) {
	suite.Run(t, &restSuite{
		makeStorer: func() store.Storer { return store.NewMapStore() },
	})
	suite.Run(t, &restSuite{
		makeStorer: func() store.Storer { return store.NewSyncStore() },
	})
}

func (s *restSuite) TestHandleGetPuppy() {
	assert := assert.New(s.T())
	puppyServer := httptest.NewServer(s.handler)
	expectedPuppy := puppy.Puppy{
		ID:     1,
		Breed:  "dog",
		Colour: "white",
		Value:  "2",
	}
	resp, err := http.Get(puppyServer.URL + "/api/puppy/1")
	assert.NoError(err)
	defer resp.Body.Close()
	actual, err := ioutil.ReadAll(resp.Body)
	assert.NoError(err)
	actualPuppy := puppy.Puppy{}
	err = json.Unmarshal(actual, &actualPuppy)
	assert.NoError(err)
	assert.Equal(200, resp.StatusCode)
	assert.Equal(expectedPuppy, actualPuppy)
}

func (s *restSuite) TestHandleGetMissingPuppy() {
	assert := assert.New(s.T())
	puppyServer := httptest.NewServer(s.handler)
	expectedPuppy := puppy.Puppy{}
	resp, err := http.Get(puppyServer.URL + "/api/puppy/non-existing")
	assert.NoError(err)
	defer resp.Body.Close()
	actual, err := ioutil.ReadAll(resp.Body)
	assert.NoError(err)
	actualPuppy := puppy.Puppy{}
	err = json.Unmarshal(actual, &actualPuppy)
	assert.NoError(err)
	assert.Equal(404, resp.StatusCode)
	assert.Equal(expectedPuppy, actualPuppy)
}

func (s *restSuite) TestHandleCreatePuppy() {
	assert := assert.New(s.T())
	puppyServer := httptest.NewServer(s.handler)
	newPuppy := puppy.Puppy{
		ID:     100,
		Breed:  "dog",
		Colour: "red",
		Value:  "20",
	}
	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(newPuppy)
	assert.NoError(err)
	resp, err := http.Post(puppyServer.URL+"/api/puppy", "application/json; charset=utf-8", b)
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)
	defer resp.Body.Close()
}

func (s *restSuite) TestHandleCreateInvalidPuppy() {
	assert := assert.New(s.T())
	puppyServer := httptest.NewServer(s.handler)
	newPuppy := puppy.Puppy{
		ID:     100,
		Breed:  "dog",
		Colour: "red",
		Value:  "-2",
	}
	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(newPuppy)
	assert.NoError(err)
	resp, err := http.Post(puppyServer.URL+"/api/puppy", "application/json; charset=utf-8", b)
	assert.NoError(err)
	defer resp.Body.Close()
	assert.Equal(http.StatusBadRequest, resp.StatusCode)
}

func (s *restSuite) TestHandleCreateBadPuppy() {
	assert := assert.New(s.T())
	puppyServer := httptest.NewServer(s.handler)
	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(`{bad json}`)
	assert.NoError(err)
	resp, err := http.Post(puppyServer.URL+"/api/puppy", "application/json; charset=utf-8", b)
	assert.NoError(err)
	defer resp.Body.Close()
	assert.Equal(http.StatusBadRequest, resp.StatusCode)
}

func (s *restSuite) TestHandleUpdatePuppy() {
	puppyServer := httptest.NewServer(s.handler)
	assert := assert.New(s.T())
	update := puppy.Puppy{
		ID:     1,
		Breed:  "dog",
		Colour: "red",
		Value:  "20",
	}
	b, err := json.Marshal(update)
	assert.NoError(err)
	req, err := http.NewRequest(http.MethodPut, puppyServer.URL+"/api/puppy", bytes.NewBuffer(b))
	assert.NoError(err)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	client := &http.Client{}
	resp, err := client.Do(req)
	assert.NoError(err)
	defer resp.Body.Close()
	actual, err := ioutil.ReadAll(resp.Body)
	assert.NoError(err)
	updatedPuppy := puppy.Puppy{}
	err = json.Unmarshal(actual, &updatedPuppy)
	assert.NoError(err)
	assert.Equal("red", updatedPuppy.Colour)
}

func (s *restSuite) TestHandleUpdateMissingPuppy() {
	assert := assert.New(s.T())
	puppyServer := httptest.NewServer(s.handler)
	updatePuppy := puppy.Puppy{
		ID:     100,
		Breed:  "dog",
		Colour: "red",
		Value:  "2",
	}
	b, err := json.Marshal(updatePuppy)
	assert.NoError(err)
	req, err := http.NewRequest(http.MethodPut, puppyServer.URL+"/api/puppy", bytes.NewBuffer(b))
	assert.NoError(err)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	client := &http.Client{}
	resp, err := client.Do(req)
	assert.NoError(err)
	defer resp.Body.Close()
	actual, err := ioutil.ReadAll(resp.Body)
	assert.NoError(err)
	updatedPuppy := puppy.Puppy{}
	err = json.Unmarshal(actual, &updatedPuppy)
	assert.NoError(err)
	assert.Equal(404, resp.StatusCode)
	assert.Equal(puppy.Puppy{}, updatedPuppy)
}

func (s *restSuite) TestHandleUpdateBadPuppy() {
	assert := assert.New(s.T())
	puppyServer := httptest.NewServer(s.handler)
	b, err := json.Marshal(`{bad json}`)
	assert.NoError(err)
	req, err := http.NewRequest(http.MethodPut, puppyServer.URL+"/api/puppy", bytes.NewBuffer(b))
	assert.NoError(err)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	client := &http.Client{}
	resp, err := client.Do(req)
	assert.NoError(err)
	defer resp.Body.Close()
	actual, err := ioutil.ReadAll(resp.Body)
	assert.NoError(err)
	updatedPuppy := puppy.Puppy{}
	err = json.Unmarshal(actual, &updatedPuppy)
	assert.NoError(err)
	assert.Equal(400, resp.StatusCode)
	assert.Equal(puppy.Puppy{}, updatedPuppy)
}

func (s *restSuite) TestDeletePuppy() {
	assert := assert.New(s.T())
	puppyServer := httptest.NewServer(s.handler)
	req, err := http.NewRequest(http.MethodDelete, puppyServer.URL+"/api/puppy/1", nil)
	assert.NoError(err)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	client := &http.Client{}
	resp, err := client.Do(req)
	assert.NoError(err)
	defer resp.Body.Close()
	assert.Equal(http.StatusNoContent, resp.StatusCode)
	resp, err = http.Get(puppyServer.URL + "/api/puppy/1")
	assert.NoError(err)
	defer resp.Body.Close()
	assert.Equal(404, resp.StatusCode)
}

func (s *restSuite) TestHandleDeleteBadPuppy() {
	assert := assert.New(s.T())
	puppyServer := httptest.NewServer(s.handler)
	req, err := http.NewRequest(http.MethodDelete, puppyServer.URL+"/api/puppy/non-existing", nil)
	assert.NoError(err)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	client := &http.Client{}
	resp, err := client.Do(req)
	assert.NoError(err)
	defer resp.Body.Close()
	assert.Equal(404, resp.StatusCode)
}

func (s *restSuite) TestHandleDeleteMissingPuppy() {
	assert := assert.New(s.T())
	puppyServer := httptest.NewServer(s.handler)
	req, err := http.NewRequest(http.MethodDelete, puppyServer.URL+"/api/puppy/78", nil)
	assert.NoError(err)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	client := &http.Client{}
	resp, err := client.Do(req)
	assert.NoError(err)
	defer resp.Body.Close()
	assert.Equal(404, resp.StatusCode)
}

func initialisePuppyStore(storer store.Storer, fileName string) error {
	puppies := []puppy.Puppy{}
	puppiesBytes := readFile(fileName)
	if err := json.Unmarshal(puppiesBytes, &puppies); err != nil {
		panic(err)
	}
	for _, puppy := range puppies {
		if err := storer.CreatePuppy(puppy); err != nil {
			return err
		}
	}
	return nil
}

func readFile(filename string) []byte {
	buff, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return buff
}
