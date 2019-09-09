package puppy_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	chi "github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	puppy "github.com/anz-bank/go-course/10_rest/runnerdave/pkg/puppy"
	store "github.com/anz-bank/go-course/10_rest/runnerdave/pkg/puppy/store"
)

var (
	rs puppy.RestStorer
)

func startAPIHandler() *chi.Mux {
	r := chi.NewRouter()
	rs = puppy.RestStorer{store.NewSyncStore()}
	newPuppy := puppy.Puppy{
		Breed:  "ULTRASURE",
		Colour: "green",
		Value:  2732.81,
	}
	if err := rs.Db.CreatePuppy(newPuppy); err != nil {
		panic(err)
	}

	r.Get(rs.GetPuppyRoute(), rs.GetPuppy)
	r.Post(rs.PostPuppyRoute(), rs.CreatePuppy)
	r.Put(rs.PutPuppyRoute(), rs.UpdatePuppy)
	r.Delete(rs.DeletePuppyRoute(), rs.DeletePuppy)

	return r
}

func TestGetPuppy(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/puppy/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := startAPIHandler()

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{
		"id": 1,
		"breed": "ULTRASURE",
		"color": "green",
		"value": 2732.81
	  }`
	actual := rr.Body.String()
	require.JSONEq(t, expected, actual, "GET Puppy was incorrect, got: %s, want: %s.",
		actual, expected)
}

func TestGetPuppyNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/puppy/200", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := startAPIHandler()

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{
		"id": 0,
		"breed": "",
		"value": 0
	  }`
	actual := rr.Body.String()
	require.JSONEq(t, expected, actual, "GET Puppy was incorrect, got: %s, want: %s.",
		actual, expected)
}

func TestPostPuppy(t *testing.T) {
	newPuppy := puppy.Puppy{
		ID:     2,
		Breed:  "EPLODE",
		Colour: "brown",
		Value:  3889.92,
	}
	body, _ := json.Marshal(newPuppy)
	req, err := http.NewRequest("POST", "/api/puppy/", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := startAPIHandler()
	handler.ServeHTTP(rr, req)

	savedPuppy, _ := rs.Db.ReadPuppy(2)
	assert.Equal(t, newPuppy, savedPuppy)
}

func TestPostInvalidPuppy(t *testing.T) {
	newPuppy := `{
		"IDE":     3,
		"Breeder":  "EPLODE",
		"Colour": "brown",
		"Value":  "a",
	}`
	body, _ := json.Marshal(newPuppy)
	req, err := http.NewRequest("POST", "/api/puppy/", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := startAPIHandler()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestPutPuppy(t *testing.T) {
	updatedPuppy := puppy.Puppy{
		ID:     1,
		Breed:  "EPLODE",
		Colour: "brown",
		Value:  3889.92,
	}
	body, _ := json.Marshal(updatedPuppy)
	req, err := http.NewRequest("PUT", "/api/puppy/1", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := startAPIHandler()
	handler.ServeHTTP(rr, req)

	savedPuppy, _ := rs.Db.ReadPuppy(1)
	assert.Equal(t, updatedPuppy, savedPuppy)
}

func TestPutInvalidPuppy(t *testing.T) {
	newPuppy := `{
		"IDE":     3,
		"Breeder":  "EPLODE",
		"Colour": "brown",
		"Value":  "a",
	}`
	body, _ := json.Marshal(newPuppy)
	req, err := http.NewRequest("PUT", "/api/puppy/2", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := startAPIHandler()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestDeletePuppy(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/api/puppy/2", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := startAPIHandler()
	handler.ServeHTTP(rr, req)

	savedPuppy, _ := rs.Db.ReadPuppy(2)
	assert.Equal(t, puppy.Puppy{}, savedPuppy)
}

func TestDeleteNotExistingPuppy(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/api/puppy/22", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := startAPIHandler()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}

func TestDeleteInvalidIdPuppy(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/api/puppy/ss", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := startAPIHandler()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}
