package puppy

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/stretchr/testify/assert"
)

func establishTestRouter() *chi.Mux {
	// create storer
	storer := NewMapStore()
	newPuppy := Puppy{ID: 1, Breed: "Jack Russell Terrier", Color: "White and Brown", Value: "1500"}
	if _, err := storer.CreatePuppy(&newPuppy); err != nil {
		panic(err)
	}

	// create rest handler
	puppyHandler := NewRestHandler(storer, lostPuppyServer)
	r := chi.NewRouter()
	r.Use(middleware.URLFormat)
	// Setup routes
	puppyHandler.SetupRoutes(r)

	return r
}

func runSubTest(t *testing.T, httpMethod string, url string, payload []byte, expected string, expectedHTTPCode int) {
	req, err := http.NewRequest(httpMethod, url, bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	// setup router
	router := establishTestRouter()
	// Setup recorder
	rr := httptest.NewRecorder()

	// Execute
	router.ServeHTTP(rr, req)

	// Get actual response
	resp := rr.Result()
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(body)

	// check for http status
	assert.Equal(t, expectedHTTPCode, resp.StatusCode)
	// check for body
	assert.Equal(t, expected, bodyString)
}

func TestRestAPI(t *testing.T) {
	tests := []struct {
		testName    string
		httpMethod  string
		url         string
		payload     []byte
		HTTPCode    int
		expectedMsg string
	}{
		{
			testName:    "Test GET Puppy by ID",
			httpMethod:  "GET",
			url:         "/api/puppy/1",
			payload:     nil,
			HTTPCode:    http.StatusOK,
			expectedMsg: "{\"id\":1,\"breed\":\"Jack Russell Terrier\",\"color\":\"White and Brown\",\"value\":\"1500\"}\n"},
		{
			testName:    "Test GET Puppy with invalid input",
			httpMethod:  "GET",
			url:         "/api/puppy/foo",
			payload:     nil,
			HTTPCode:    http.StatusBadRequest,
			expectedMsg: "Bad Request: Invalid input, ensure ID and JSON are valid\n"},
		{
			testName:    "Test GET invalid Puppy",
			httpMethod:  "GET",
			url:         "/api/puppy/100",
			payload:     nil,
			HTTPCode:    http.StatusNotFound,
			expectedMsg: "Not Found: The puppy ID does not exist\n"},
		{
			testName:    "Test POST Puppy",
			httpMethod:  "POST",
			url:         "/api/puppy/",
			payload:     []byte(`{"Breed": "German Shepperd", "Color": "Brown", "Value": "2000"}`),
			HTTPCode:    http.StatusCreated,
			expectedMsg: "{\"id\":2,\"breed\":\"German Shepperd\",\"color\":\"Brown\",\"value\":\"2000\"}\n"},
		{
			testName:    "Test POST with error data",
			httpMethod:  "POST",
			url:         "/api/puppy/",
			payload:     []byte(`{"Breed": "German Shepperd", "Color": "Brown", "Value": "-2000"}`),
			HTTPCode:    http.StatusBadRequest,
			expectedMsg: "Bad Request: Invalid input, ensure ID and JSON are valid\n"},
		{
			testName:    "Test POST with decode error",
			httpMethod:  "POST",
			url:         "/api/puppy/",
			payload:     []byte(`{"What: "I don't know", "When": "Now", "Where": "Somewhere"}`),
			HTTPCode:    http.StatusBadRequest,
			expectedMsg: "Bad Request: Invalid input, ensure ID and JSON are valid\n"},
		{
			testName:    "Test PUT puppy by ID",
			httpMethod:  "PUT",
			url:         "/api/puppy/1",
			payload:     []byte(`{"id": 1, "Breed": "German Shepperd", "Color": "Brown", "Value": "8888"}`),
			HTTPCode:    http.StatusOK,
			expectedMsg: "\"Puppy is updated successfully\"\n"},
		{
			testName:    "Test PUT puppy with invalid payload",
			httpMethod:  "PUT",
			url:         "/api/puppy/1",
			payload:     []byte(`{INVALID JSON}`),
			HTTPCode:    http.StatusUnprocessableEntity,
			expectedMsg: "Unprocessable Entity: Invalid input, ensure ID and JSON are valid\n"},
		{
			testName:    "Test PUT puppy with invalid id",
			httpMethod:  "PUT",
			url:         "/api/puppy/100",
			payload:     []byte(`{"id": 100, "Breed": "German Shepperd", "Color": "Brown", "Value": "2000"}`),
			HTTPCode:    http.StatusNotFound,
			expectedMsg: "Not Found: The puppy ID does not exist\n"},
		{
			testName:    "Test PUT with invalid id",
			httpMethod:  "PUT",
			url:         "/api/puppy/foo",
			payload:     []byte(`{"Breed": "German Shepperd", "Color": "Brown", "Value": "2000"}`),
			HTTPCode:    http.StatusBadRequest,
			expectedMsg: "Bad Request: Invalid input, ensure ID and JSON are valid\n"},
		{
			testName:    "Test PUT with negative puppy value",
			httpMethod:  "PUT",
			url:         "/api/puppy/1",
			payload:     []byte(`{"Breed": "German Shepperd", "Color": "Brown", "Value": "-2000"}`),
			HTTPCode:    http.StatusBadRequest,
			expectedMsg: "Bad Request: Invalid input, ensure ID and JSON are valid\n"},
		{
			testName:    "Test PUT with non-int puppy value",
			httpMethod:  "PUT",
			url:         "/api/puppy/1",
			payload:     []byte(`{"Breed": "German Shepperd", "Color": "Brown", "Value": "blah"}`),
			HTTPCode:    http.StatusBadRequest,
			expectedMsg: "Bad Request: Invalid input, ensure ID and JSON are valid\n"},
		{
			testName:    "Test DELETE puppy by ID",
			httpMethod:  "DELETE",
			url:         "/api/puppy/1",
			payload:     nil,
			HTTPCode:    http.StatusOK,
			expectedMsg: "\"Puppy is deleted successfully\"\n"},
		{
			testName:    "Test DELETE puppy with non existent id",
			httpMethod:  "DELETE",
			url:         "/api/puppy/100",
			payload:     nil,
			HTTPCode:    http.StatusNotFound,
			expectedMsg: "Not Found: Invalid input, ensure ID and JSON are valid\n"},
		{
			testName:    "Test DELETE puppy with invalid id",
			httpMethod:  "DELETE",
			url:         "/api/puppy/foo",
			payload:     nil,
			HTTPCode:    http.StatusBadRequest,
			expectedMsg: "Bad Request: Invalid input, ensure ID and JSON are valid\n"},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.testName, func(t *testing.T) {
			runSubTest(t, tc.httpMethod, tc.url, tc.payload, tc.expectedMsg, tc.HTTPCode)
		})
	}
}

// Special thanks to Daniel Cantos for helping/teaching me mocking api responses
// which pushes my test coverage from 98->100% :)
// the code below overwrite default http client
// this covers api post call inside a goroutine NotifyLostPuppy()
type rtFunc func(req *http.Request) (*http.Response, error)

func (r rtFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return r(req)
}

func badResponseClient() *http.Client {
	return &http.Client{
		Transport: rtFunc(func(req *http.Request) (*http.Response, error) {
			return nil, errors.New("Something went wrong, the request was not successful")
		}),
	}
}

func serverMock(status int) *httptest.Server {
	handler := http.NewServeMux()
	handler.HandleFunc("/path", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
	})
	return httptest.NewServer(handler)
}

func TestNotifyLostPuppyWith201Response(t *testing.T) {
	storer := NewMapStore()
	server := serverMock(http.StatusCreated)
	ph := NewRestHandler(storer, server.URL)
	defer server.Close()
	client := server.Client()
	http.DefaultClient = client
	ph.notifyLostPuppy(1)
}

func TestNotifyLostPuppyWith500Response(t *testing.T) {
	storer := NewMapStore()
	ph := NewRestHandler(storer, lostPuppyServer)
	http.DefaultClient = badResponseClient()
	ph.notifyLostPuppy(111)
}
