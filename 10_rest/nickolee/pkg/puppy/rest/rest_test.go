package rest

import (
	"bytes"
	"errors"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/anz-bank/go-course/10_rest/nickolee/pkg/puppy"
	"github.com/anz-bank/go-course/10_rest/nickolee/pkg/puppy/store"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func TestHandleStorerError(t *testing.T) {
	wr := httptest.NewRecorder()
	err := errors.New("hello")
	handleStorerError(wr, err)
	resp := wr.Result() // can treat as if is the response from the handler
	require.Equal(t, 500, resp.StatusCode)

	raw, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	require.NoError(t, err)
	require.Equal(t, "500: Internal Server Error\n", string(raw))
}

func createStore(n int) puppy.Storer {
	switch n {
	case 0:
		return store.NewMapStore()
	default:
		return store.NewSyncStore()
	}
}

// a setup function that returns a preconfigured router that our test client can use
// note that we passed it a *testing.T
func createTestRouter(t *testing.T) *chi.Mux {
	// randomly create a SyncStore or a MapStore for tests
	storer := createStore(rand.Intn(2)) // generates random binary number 0 or 1

	// create new PuppyHandlerAndStorer with some seed data
	phs := NewPuppyHandlerAndStorer(storer)
	seedPuppy := puppy.Puppy{Breed: "Extremely Rare Golden Retriever",
		Colour: "An extremely rare shade of violet", Value: 1000000} // so you know it's going to cost you :)
	_, err := phs.Storage.CreatePuppy(&seedPuppy)
	require.NoError(t, err)

	// create chi test router
	r := chi.NewRouter()
	r.Use(middleware.URLFormat)

	// Setup routes
	SetupRoutes(r, *phs)

	return r
}

func setupTest(t *testing.T) (*httptest.Server, *http.Client) {
	r := createTestRouter(t)
	s := httptest.NewServer(r) // Wrapping your handler in a server struct
	client := s.Client()       // Creating a client that can talk to that server
	return s, client
}

func TestGetPuppy(t *testing.T) {
	// prepare test data
	testCases := []struct {
		testName     string
		url          string
		HTTPCode     int
		expectedBody string
	}{
		{
			testName: "Test GET Puppy by ID",
			url:      "/api/puppy/1",
			HTTPCode: http.StatusOK,
			expectedBody: "{\"id\":1,\"breed\":\"Extremely Rare Golden Retriever\",\"colour\":\"An extremely" +
				" rare shade of violet\",\"value\":1000000}\n"},
		{
			testName:     "Test GET Puppy on invalid endpoint",
			url:          "/api/puppy/invalid",
			HTTPCode:     http.StatusBadRequest,
			expectedBody: "Bad Request: Invalid input, ensure ID is valid\n"},
		{
			testName:     "Test GET Puppy on nonexistent Puppy",
			url:          "/api/puppy/300",
			HTTPCode:     http.StatusNotFound,
			expectedBody: "PuppyStoreError 404: Sorry puppy with ID 300 does not exist\n"},
	}

	// setup test server, router and client + starts server
	// This guy basically is like a one-use server that we stand up then tear down for each test case
	// It only exists for executing one test case
	ts, client := setupTest(t) // somehow this doesn't look like it but setupTest() actually also starts the server
	defer ts.Close()

	// run tests by looping through test cases
	for _, tc := range testCases {
		tc := tc // prevents scopelint error
		t.Run(tc.testName, func(t *testing.T) {
			// Call endpoint
			resp, err := client.Get(ts.URL + tc.url) // ts.URL is the server's base URL
			require.NoError(t, err)
			require.NotNil(t, resp)

			// Check response
			assert.Equal(t, tc.HTTPCode, resp.StatusCode) // compare actual vs expected response code
			defer resp.Body.Close()                       // ensure you close the read on the next line
			body, err := ioutil.ReadAll(resp.Body)        // check body of resp
			require.NoError(t, err)
			bodyString := string(body)                   // all this while we just had a stream of bytes. Now we get the text
			assert.Equal(t, tc.expectedBody, bodyString) // check body - whether expected == actual
		})
	}
}

func TestPostPuppy(t *testing.T) {
	// prepare test data
	testCases := []struct {
		testName     string
		url          string
		payload      []byte
		HTTPCode     int
		expectedBody string
	}{
		{
			testName:     "Test POST Puppy",
			url:          "/api/puppy/",
			payload:      []byte(`{"Breed": "Arcanine", "Colour": "Brown", "Value": 7000}`),
			HTTPCode:     http.StatusCreated,
			expectedBody: "{\"id\":2,\"breed\":\"Arcanine\",\"colour\":\"Brown\",\"value\":7000}\n"},
		{
			testName:     "Test POST with negative value",
			url:          "/api/puppy/",
			payload:      []byte(`{"Breed": "Arcanine", "Colour": "Brown", "Value": -7000}`),
			HTTPCode:     http.StatusBadRequest,
			expectedBody: "PuppyStoreError 400: Sorry puppy value cannot be negative. The dog has to be worth something :)\n"},
		{
			testName:     "Test POST with JSON decode error",
			url:          "/api/puppy/",
			payload:      []byte(`{"What: "I don't know", "When": "Now", "Where": "Somewhere"}`),
			HTTPCode:     http.StatusBadRequest,
			expectedBody: "Bad Request: Invalid input, ensure JSON is valid\n"},
	}

	// setup test server, router and client + starts server
	ts, client := setupTest(t) // somehow this doesn't look like it but setupTest() actually also starts the server
	defer ts.Close()

	// run tests by looping through test cases
	for _, tc := range testCases {
		tc := tc // prevents scopelint error
		t.Run(tc.testName, func(t *testing.T) {
			// Call endpoint (ts.URL is the server's base URL)
			resp, err := client.Post(ts.URL+tc.url, "application/json", bytes.NewBuffer(tc.payload))
			require.NoError(t, err)
			require.NotNil(t, resp)

			// Check response
			assert.Equal(t, tc.HTTPCode, resp.StatusCode) // compare actual vs expected response code
			defer resp.Body.Close()                       // ensure you close the read on the next line
			body, err := ioutil.ReadAll(resp.Body)        // check body of resp
			require.NoError(t, err)
			bodyString := string(body) // all this while we just had a stream of bytes. Now we get the text
			// fmt.Printf("Test Case %s ran: %v\n", tc.testName, bodyString)
			assert.Equal(t, tc.expectedBody, bodyString) // check body - whether expected == actual
		})
	}
}

func TestPutPuppy(t *testing.T) {
	// prepare test data
	testCases := []struct {
		testName     string
		url          string
		payload      []byte
		HTTPCode     int
		expectedBody string
	}{
		{
			testName:     "Test PUT puppy by ID",
			url:          "/api/puppy/1",
			payload:      []byte(`{"id": 1, "Breed": "Vulpix", "Colour": "Brown-ish", "Value": 2}`),
			HTTPCode:     http.StatusCreated,
			expectedBody: "{\"id\":1,\"breed\":\"Vulpix\",\"colour\":\"Brown-ish\",\"value\":2}\n"},
		{
			testName:     "Test PUT puppy with invalid payload",
			url:          "/api/puppy/1",
			payload:      []byte(`{INVALID JSON}`),
			HTTPCode:     http.StatusUnprocessableEntity,
			expectedBody: "Unprocessable Entity: Invalid input, ensure JSON is valid\n"},
		{
			testName:     "Test PUT puppy with out of range id",
			url:          "/api/puppy/300",
			payload:      []byte(`{"id": 300, "Breed": "Pikachu", "Colour": "Yellow", "Value": 2100}`),
			HTTPCode:     http.StatusNotFound,
			expectedBody: "PuppyStoreError 404: Sorry puppy with ID 300 does not exist\n"},
		{
			testName:     "Test PUT with invalid id",
			url:          "/api/puppy/foo",
			payload:      []byte(`{"Breed": "Liverbird", "Color": "Red", "Value": 20000}`),
			HTTPCode:     http.StatusBadRequest,
			expectedBody: "Unprocessable Entity: Invalid input, ensure ID is valid\n"},
		{
			testName:     "Test PUT with non-int puppy value",
			url:          "/api/puppy/1",
			payload:      []byte(`{"Breed": "T-Rex", "Colour": "Green-ish", "Value": "blah"}`),
			HTTPCode:     http.StatusUnprocessableEntity,
			expectedBody: "Unprocessable Entity: Invalid input, ensure JSON is valid\n"},
		{
			testName:     "Test PUT with negative puppy value",
			url:          "/api/puppy/1",
			payload:      []byte(`{"Breed": "Sphinx", "Color": "Golden", "Value": -500000}`),
			HTTPCode:     http.StatusBadRequest,
			expectedBody: "PuppyStoreError 400: Sorry puppy value cannot be negative. The dog has to be worth something :)\n"},
	}

	// setup test server, router and client + starts server
	ts, client := setupTest(t) // somehow this doesn't look like it but setupTest() actually also starts the server
	defer ts.Close()

	// run tests by looping through test cases
	for _, tc := range testCases {
		tc := tc // prevents scopelint error
		t.Run(tc.testName, func(t *testing.T) {
			// setup request
			req, err := http.NewRequest("PUT", ts.URL+tc.url, bytes.NewBuffer(tc.payload))
			require.NoError(t, err)

			// Call endpoint
			resp, err := client.Do(req) // ts.URL is the server's base URL
			require.NoError(t, err)
			require.NotNil(t, resp)

			// Check response
			assert.Equal(t, tc.HTTPCode, resp.StatusCode) // compare actual vs expected response code
			defer resp.Body.Close()                       // ensure you close the read on the next line
			body, err := ioutil.ReadAll(resp.Body)        // check body of resp
			require.NoError(t, err)
			bodyString := string(body)                   // all this while we just had a stream of bytes. Now we get the text
			assert.Equal(t, tc.expectedBody, bodyString) // check body - whether expected == actual
		})
	}
}

func TestDeletePuppy(t *testing.T) {
	// prepare test data
	testCases := []struct {
		testName     string
		url          string
		HTTPCode     int
		expectedBody string
	}{
		{
			testName:     "Test DELETE puppy by ID",
			url:          "/api/puppy/1",
			HTTPCode:     http.StatusOK,
			expectedBody: "\"Puppy successfully deleted\"\n"},
		{
			testName:     "Test DELETE puppy with non-existent id",
			url:          "/api/puppy/300",
			HTTPCode:     http.StatusNotFound,
			expectedBody: "PuppyStoreError 404: Sorry puppy with ID 300 does not exist\n"},
		{
			testName:     "Test DELETE puppy with invalid id",
			url:          "/api/puppy/invalid",
			HTTPCode:     http.StatusBadRequest,
			expectedBody: "Bad Request: Invalid input, ensure ID is valid\n"},
	}

	// setup test server, router and client + starts server
	ts, client := setupTest(t) // somehow this doesn't look like it but setupTest() actually also starts the server
	defer ts.Close()

	// run tests by looping through test cases
	for _, tc := range testCases {
		tc := tc // prevents scopelint error
		t.Run(tc.testName, func(t *testing.T) {
			// setup request
			req, err := http.NewRequest("DELETE", ts.URL+tc.url, nil)
			require.NoError(t, err)

			// Call endpoint
			resp, err := client.Do(req) // ts.URL is the server's base URL
			require.NoError(t, err)
			require.NotNil(t, resp)

			// Check response
			assert.Equal(t, tc.HTTPCode, resp.StatusCode) // compare actual vs expected response code
			defer resp.Body.Close()                       // ensure you close the read on the next line
			body, err := ioutil.ReadAll(resp.Body)        // check body of resp
			require.NoError(t, err)
			bodyString := string(body)                   // all this while we just had a stream of bytes. Now we get the text
			assert.Equal(t, tc.expectedBody, bodyString) // check body - whether expected == actual
		})
	}
}
