package puppy_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	puppy "github.com/anz-bank/go-course/10_rest/runnerdave/pkg/puppy"
	store "github.com/anz-bank/go-course/10_rest/runnerdave/pkg/puppy/store"
)

var (
	rs puppy.RestServer
)

func TestGetPuppy(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/puppy/1", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	newPuppy := puppy.Puppy{
		Breed:  "ULTRASURE",
		Colour: "green",
		Value:  2732.81,
	}
	db := store.NewSyncStore()
	err = db.CreatePuppy(newPuppy)
	require.NoError(t, err)

	handler := rs.SetupRoutes(db)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	expected := `{
		"id": 1,
		"breed": "ULTRASURE",
		"color": "green",
		"value": 2732.81
	  }`
	assert.JSONEq(t, expected, rr.Body.String())
}

func TestGetPuppyNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/puppy/200", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := rs.SetupRoutes(store.NewSyncStore())
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
	assert.Equal(t, "", rr.Body.String())
}

func TestGetPuppyBadID(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/puppy/bad", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := rs.SetupRoutes(store.NewSyncStore())
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, "", rr.Body.String())
}

func TestPostPuppy(t *testing.T) {
	newPuppy := puppy.Puppy{
		ID:     1,
		Breed:  "EPLODE",
		Colour: "brown",
		Value:  3889.92,
	}
	body, err := json.Marshal(newPuppy)
	require.NoError(t, err)

	req, err := http.NewRequest("POST", "/api/puppy/", bytes.NewBuffer(body))
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := rs.SetupRoutes(store.NewSyncStore())
	handler.ServeHTTP(rr, req)

	savedPuppy, err := rs.DB.ReadPuppy(1)
	require.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Equal(t, newPuppy, savedPuppy)
}

func TestPostInvalidPuppy(t *testing.T) {
	newPuppy := `{
		"IDE":     3,
		"Breeder":  "EPLODE",
		"Colour": "brown",
		"Value":  "a",
	}`
	body, err := json.Marshal(newPuppy)
	require.NoError(t, err)
	req, err := http.NewRequest("POST", "/api/puppy/", bytes.NewBuffer(body))
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := rs.SetupRoutes(store.NewSyncStore())
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestPostInvalidValuePuppy(t *testing.T) {
	newPuppy := puppy.Puppy{
		ID:     1,
		Breed:  "EPLODE",
		Colour: "brown",
		Value:  -3889.92,
	}
	body, err := json.Marshal(newPuppy)
	require.NoError(t, err)
	req, err := http.NewRequest("POST", "/api/puppy/", bytes.NewBuffer(body))
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := rs.SetupRoutes(store.NewSyncStore())
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestPutPuppy(t *testing.T) {
	updatedPuppy := puppy.Puppy{
		ID:     1,
		Breed:  "EPLODE",
		Colour: "brown",
		Value:  3889.92,
	}
	body, err := json.Marshal(updatedPuppy)
	require.NoError(t, err)
	req, err := http.NewRequest("PUT", "/api/puppy/1", bytes.NewBuffer(body))
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	newPuppy := puppy.Puppy{
		Breed:  "ULTRASURE",
		Colour: "green",
		Value:  2732.81,
	}
	db := store.NewSyncStore()
	err = db.CreatePuppy(newPuppy)
	require.NoError(t, err)

	handler := rs.SetupRoutes(db)
	handler.ServeHTTP(rr, req)

	savedPuppy, err := rs.DB.ReadPuppy(1)
	require.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, rr.Code)
	assert.Equal(t, updatedPuppy, savedPuppy)
}

func TestPutInvalidValuePuppy(t *testing.T) {
	updatedPuppy := puppy.Puppy{
		ID:     1,
		Breed:  "EPLODE",
		Colour: "brown",
		Value:  -3889.92,
	}
	body, err := json.Marshal(updatedPuppy)
	require.NoError(t, err)
	req, err := http.NewRequest("PUT", "/api/puppy/1", bytes.NewBuffer(body))
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	db := store.NewSyncStore()

	handler := rs.SetupRoutes(db)
	handler.ServeHTTP(rr, req)

	require.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestPutInvalidPuppy(t *testing.T) {
	newPuppy := `{
		"IDE":     3,
		"Breeder":  "EPLODE",
		"Colour": "brown",
		"Value":  "a",
	}`
	body, err := json.Marshal(newPuppy)
	require.NoError(t, err)
	req, err := http.NewRequest("PUT", "/api/puppy/2", bytes.NewBuffer(body))
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := rs.SetupRoutes(store.NewSyncStore())
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestPutInvalidPuppyID(t *testing.T) {
	newPuppy := `{
		"IDE":     3,
		"Breeder":  "EPLODE",
		"Colour": "brown",
		"Value":  "12.4",
	}`
	body, err := json.Marshal(newPuppy)
	require.NoError(t, err)
	req, err := http.NewRequest("PUT", "/api/puppy/bad", bytes.NewBuffer(body))
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := rs.SetupRoutes(store.NewSyncStore())
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestDeletePuppy(t *testing.T) {
	rr := httptest.NewRecorder()
	newPuppy := puppy.Puppy{
		Breed:  "ULTRASURE",
		Colour: "green",
		Value:  2732.81,
	}
	db := store.NewSyncStore()
	err := db.CreatePuppy(newPuppy)
	require.NoError(t, err)

	req, err := http.NewRequest("DELETE", "/api/puppy/1", nil)
	require.NoError(t, err)

	handler := rs.SetupRoutes(db)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)

	savedPuppy, err := rs.DB.ReadPuppy(1)
	require.Error(t, err)
	assert.Equal(t, puppy.Puppy{}, savedPuppy)
}

func TestDeleteNotExistingPuppy(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/api/puppy/22", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := rs.SetupRoutes(store.NewSyncStore())
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestDeleteInvalidIDPuppy(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/api/puppy/ss", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := rs.SetupRoutes(store.NewSyncStore())
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}
