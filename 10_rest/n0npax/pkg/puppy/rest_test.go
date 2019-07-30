package puppy_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	tassert "github.com/stretchr/testify/assert"

	puppy "github.com/anz-bank/go-course/10_rest/n0npax/pkg/puppy"
	store "github.com/anz-bank/go-course/10_rest/n0npax/pkg/puppy/store"
	"github.com/stretchr/testify/suite"
)

type storerSuite struct {
	suite.Suite
	store puppy.Storer
}

func (s *storerSuite) SetupTest() {
	switch s.store.(type) {
	case *store.SyncStore:
		s.store = store.NewSyncStore()
	case *store.MemStore:
		s.store = store.NewMemStore()
	default:
		panic("Unrecognised storer implementation")
	}

	for i := 0; i < 5; i++ {
		p := puppy.Puppy{
			Breed:  "Type A",
			Colour: "Grey",
			Value:  100 * i,
		}
		_, err := s.store.CreatePuppy(&p)
		if err != nil {
			panic(err)
		}
	}
}

func TestStorer(t *testing.T) {
	suite.Run(t, &storerSuite{store: store.NewMemStore()})
	suite.Run(t, &storerSuite{store: store.NewSyncStore()})
}

func sendRequest(r http.Handler, method, path string, payload io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, payload)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func getReq(r http.Handler, path string, payload io.Reader) *httptest.ResponseRecorder {
	return sendRequest(r, "GET", path, payload)
}
func putReq(r http.Handler, path string, payload io.Reader) *httptest.ResponseRecorder {
	return sendRequest(r, "PUT", path, payload)
}
func postReq(r http.Handler, path string, payload io.Reader) *httptest.ResponseRecorder {
	return sendRequest(r, "POST", path, payload)
}
func deleteReq(r http.Handler, path string, payload io.Reader) *httptest.ResponseRecorder {
	return sendRequest(r, "DELETE", path, payload)
}

// GET
func (s *storerSuite) TestGet() {
	s.Run("OK", s.GetOK)
	s.Run("NotFound", s.GetNotFound)
	s.Run("CorruptedID", s.GetCorruptedID)
}

func (s *storerSuite) GetNotFound() {
	assert := tassert.New(s.T())
	router := puppy.RestBackend(s.store)
	w := getReq(router, "/api/puppy/1000", nil)
	assert.Equal(http.StatusNotFound, w.Code)
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)

	value, exists := response["message"]
	assert.Nil(err)
	assert.True(exists)
	assert.Equal("puppy with ID (1000) not found", value)
}

func (s *storerSuite) GetCorruptedID() {
	assert := tassert.New(s.T())
	router := puppy.RestBackend(s.store)
	w := getReq(router, "/api/puppy/0x01", nil)
	assert.Equal(puppy.ErrCodeInternal, w.Code)
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)

	value, exists := response["message"]
	assert.Nil(err)
	assert.True(exists)
	assert.Equal(`strconv.Atoi: parsing "0x01": invalid syntax`, value)
}

func (s *storerSuite) GetOK() {
	assert := tassert.New(s.T())
	router := puppy.RestBackend(s.store)
	var response puppy.Puppy

	for id := 0; id < 5; id++ {
		w := getReq(router, fmt.Sprintf("/api/puppy/%d", id), nil)
		assert.Equal(http.StatusOK, w.Code)

		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.Nil(err)
		assert.Equal(id*100, response.Value)
		assert.Equal(id, response.ID)
	}
}

// POST
func (s *storerSuite) TestPost() {
	s.Run("OK", s.PostOK)
	s.Run("CorruptedID", s.PostCorrupted)
	s.Run("BadValue", s.PostBadValue)
}

func (s *storerSuite) PostCorrupted() {
	assert := tassert.New(s.T())
	router := puppy.RestBackend(s.store)
	w := postReq(router, "/api/puppy/", strings.NewReader("broken payload"))
	assert.Equal(400, w.Code)
}

func (s *storerSuite) PostOK() {
	assert := tassert.New(s.T())
	router := puppy.RestBackend(s.store)
	p := puppy.Puppy{Value: 71, Colour: "red"}
	b, err := json.Marshal(p)
	assert.Nil(err)
	payload := string(b)
	w := postReq(router, "/api/puppy/", strings.NewReader(payload))
	assert.Equal(201, w.Code)
	var response map[string]int

	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(err)
	_, ok := response["id"]
	assert.True(ok)
}

func (s *storerSuite) PostBadValue() {
	assert := tassert.New(s.T())
	router := puppy.RestBackend(s.store)
	p := puppy.Puppy{Value: -44, Colour: "red"}
	b, err := json.Marshal(p)

	assert.Nil(err)
	payload := string(b)
	w := postReq(router, "/api/puppy/", strings.NewReader(payload))
	assert.Equal(400, w.Code)
}

// Put
func (s *storerSuite) TestPut() {
	s.Run("OK", s.PutOK)
	s.Run("BadID", s.PutBadID)
	s.Run("CorruptedID", s.PutCorrupted)
	s.Run("BadValue", s.PutBadValue)
}

func (s *storerSuite) PutCorrupted() {
	assert := tassert.New(s.T())
	router := puppy.RestBackend(s.store)
	w := putReq(router, "/api/puppy/0", strings.NewReader("broken payload"))
	assert.Equal(400, w.Code)
}

func (s *storerSuite) PutBadID() {
	assert := tassert.New(s.T())
	router := puppy.RestBackend(s.store)
	w := putReq(router, "/api/puppy/0x01", strings.NewReader("broken payload"))
	assert.Equal(500, w.Code)
}

func (s *storerSuite) PutOK() {
	assert := tassert.New(s.T())
	router := puppy.RestBackend(s.store)
	p := puppy.Puppy{Value: 71, Colour: "red"}
	b, err := json.Marshal(p)

	assert.Nil(err)
	payload := string(b)
	w := putReq(router, "/api/puppy/0", strings.NewReader(payload))
	assert.Equal(204, w.Code)
}

func (s *storerSuite) PutBadValue() {
	assert := tassert.New(s.T())
	router := puppy.RestBackend(s.store)
	p := puppy.Puppy{Value: -44, Colour: "red"}
	b, err := json.Marshal(p)

	assert.Nil(err)
	payload := string(b)
	w := putReq(router, "/api/puppy/0", strings.NewReader(payload))
	assert.Equal(400, w.Code)
}

// Delete
func (s *storerSuite) TestDelete() {
	s.Run("OK", s.DeleteOK)
	s.Run("BadID", s.DeleteBadID)
	s.Run("NonExisting", s.DeleteNotExisting)
}

func (s *storerSuite) DeleteOK() {
	assert := tassert.New(s.T())
	router := puppy.RestBackend(s.store)
	w := deleteReq(router, "/api/puppy/0", nil)
	assert.Equal(204, w.Code)
}

func (s *storerSuite) DeleteBadID() {
	assert := tassert.New(s.T())
	router := puppy.RestBackend(s.store)
	w := deleteReq(router, "/api/puppy/0x01", nil)
	assert.Equal(500, w.Code)
}

func (s *storerSuite) DeleteNotExisting() {
	assert := tassert.New(s.T())
	router := puppy.RestBackend(s.store)
	w := deleteReq(router, "/api/puppy/1000", nil)
	assert.Equal(404, w.Code)
}
