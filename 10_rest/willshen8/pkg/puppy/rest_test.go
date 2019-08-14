package puppy

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/stretchr/testify/assert"
)

func establishRouter() *chi.Mux {
	// create storer
	storer := NewMapStore()
	newPuppy := Puppy{ID: 1, Breed: "Jack Russell Terrier", Color: "White and Brown", Value: "1500"}
	if _, err := storer.CreatePuppy(&newPuppy); err != nil {
		panic(err)
	}

	// create rest handler
	puppyHandler := NewRestHandler(storer)
	r := chi.NewRouter()
	r.Use(middleware.URLFormat)
	// Setup routes
	SetupRoutes(r, *puppyHandler)

	return r
}

// Enumerating test cases made life easier :)
func TestRestAPI(t *testing.T) {
	tests := []struct {
		testName   string
		httpMethod string
		url        string
		payload    []byte
		expected   string
	}{
		{
			testName:   "Test GET Puppy by ID",
			httpMethod: "GET",
			url:        "/api/puppy/1",
			payload:    nil,
			expected:   "{\"id\":1,\"breed\":\"Jack Russell Terrier\",\"color\":\"White and Brown\",\"value\":\"1500\"}\n"},
		{
			testName:   "Test GET invalid Puppy",
			httpMethod: "GET",
			url:        "/api/puppy/100",
			payload:    nil,
			expected:   "{\"message\":\"Puppy ID can not be found, read operation failed.\",\"code\":1002}\n"},
		{
			testName:   "Test POST Puppy",
			httpMethod: "POST",
			url:        "/api/puppy/",
			payload:    []byte(`{"Breed": "German Shepperd", "Color": "Brown", "Value": "2000"}`),
			expected:   "{\"id\":2,\"breed\":\"German Shepperd\",\"color\":\"Brown\",\"value\":\"2000\"}\n"},
		{
			testName:   "Test POST with error",
			httpMethod: "POST",
			url:        "/api/puppy/",
			payload:    []byte(`{"Breed": "German Shepperd", "Color": "Brown", "Value": "-2000"}`),
			expected:   "{\"message\":\"Puppy value can't be less than 0.\",\"code\":1001}\n"},
		{
			testName:   "Test POST with decode error",
			httpMethod: "POST",
			url:        "/api/puppy/",
			payload:    []byte(`{"What: "I don't know", "When": "Now", "Where": "Somewhere"}`),
			expected:   "{\"Offset\":10}\n{\"message\":\"Unrecongised puppy value.\",\"code\":1000}\n"},
		{
			testName:   "Test PUT puppy by ID",
			httpMethod: "PUT",
			url:        "/api/puppy/1",
			payload:    []byte(`{"id": 1, Breed": "German Shepperd", "Color": "Brown", "Value": "2000"}`),
			expected:   "true\n"},
		{
			testName:   "Test PUT puppy with error",
			httpMethod: "PUT",
			url:        "/api/puppy/100",
			payload:    []byte(`{"id": 100, Breed": "German Shepperd", "Color": "Brown", "Value": "2000"}`),
			expected:   "{\"message\":\"Puppy ID can not be found, update operation failed.\",\"code\":1002}\n"},
		{
			testName:   "Test DELETE puppy by ID",
			httpMethod: "DELETE",
			url:        "/api/puppy/1",
			payload:    nil,
			expected:   "true\n"},
		{
			testName:   "Test DELETE puppy with error",
			httpMethod: "DELETE",
			url:        "/api/puppy/100",
			payload:    nil,
			expected:   "{\"message\":\"Puppy ID can not be found, delete operation failed.\",\"code\":1002}\n"},
	}
	for _, tc := range tests {
		req, err := http.NewRequest(tc.httpMethod, tc.url, bytes.NewBuffer(tc.payload))
		if err != nil {
			t.Fatal(err)
		}

		//setup router
		router := establishRouter()
		//Setup recorder
		rr := httptest.NewRecorder()

		// Execute
		router.ServeHTTP(rr, req)

		// Get actual response
		resp := rr.Result()

		body, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(body)
		defer resp.Body.Close()

		//check for http status
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		//check for body
		assert.Equal(t, tc.expected, bodyString)
	}
}
