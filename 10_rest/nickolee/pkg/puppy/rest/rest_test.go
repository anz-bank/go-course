package rest

import (
	"bytes"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/anz-bank/go-course/10_rest/nickolee/pkg/puppy"
	"github.com/anz-bank/go-course/10_rest/nickolee/pkg/puppy/store"
	"github.com/stretchr/testify/assert"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func createStore(n int) puppy.Storer {
	switch n {
	case 0:
		return store.NewMapStore()
	default:
		return store.NewSyncStore()
	}
}

// a setup function that returns a preconfigured router that our test client can use
func createTestRouter() *chi.Mux {
	// randomly create a SyncStore or a MapStore for tests
	storer := createStore(rand.Intn(2)) // generates random binary number 0 or 1

	// create new PuppyHandlerAndStorer with some seed data
	phs := NewPuppyHandlerAndStorer(storer)
	seedPuppy := puppy.Puppy{Breed: "Extremely Rare Golden Retriever",
		Colour: "An extremely rare shade of violet", Value: 1000000} // so you know it's going to cost you :)
	_, err := phs.Storage.CreatePuppy(&seedPuppy)
	if err != nil {
		panic(err)
	}
	// create chi test router
	r := chi.NewRouter()
	r.Use(middleware.URLFormat)

	// Setup routes
	SetupRoutes(r, *phs)

	return r
}

// creates a function that can act as a test client to run each test case
func runTestCase(t *testing.T, httpMethod string, url string, payload []byte, expected string, expectedHTTPCode int) {
	// create mock http request
	req, err := http.NewRequest(httpMethod, url, bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	// setup router
	router := createTestRouter()

	// Setup response recorder that records the test response from the server for later inspection in tests
	rr := httptest.NewRecorder()

	/* Start test server - it needs a req and a resp which we have created above
	   This guy basically is like a one-use server that we stand up then tear down
	   It only exists for executing one test case */
	router.ServeHTTP(rr, req)

	// Get actual response from one time user server
	resp := rr.Result()                  // grab the response from the resp recorder
	defer resp.Body.Close()              // ensure you close the read on the next line
	body, _ := ioutil.ReadAll(resp.Body) // read the body
	bodyString := string(body)           // all this while we just had a stream of bytes. Now we get the text

	// check http status - whether expected == actual
	assert.Equal(t, expectedHTTPCode, resp.StatusCode)
	// check body - whether expected == actual
	assert.Equal(t, expected, bodyString)
}

func TestRestAPI(t *testing.T) {
	testCases := []struct {
		// all this is the input for the runTestCase func we created above (the test client)
		testName    string
		httpMethod  string
		url         string
		payload     []byte
		HTTPCode    int
		expectedMsg string
	}{
		{
			testName:   "Test GET Puppy by ID",
			httpMethod: "GET",
			url:        "/api/puppy/1",
			payload:    nil,
			HTTPCode:   http.StatusOK,
			expectedMsg: "{\"id\":1,\"breed\":\"Extremely Rare Golden Retriever\",\"colour\":\"An extremely" +
				" rare shade of violet\",\"value\":1000000}\n"},
		{
			testName:    "Test GET Puppy on invalid endpoint",
			httpMethod:  "GET",
			url:         "/api/puppy/invalid",
			payload:     nil,
			HTTPCode:    http.StatusBadRequest,
			expectedMsg: "Bad Request: Invalid input, ensure ID and JSON are valid\n"},
		{
			testName:    "Test GET Puppy on nonexistent Puppy",
			httpMethod:  "GET",
			url:         "/api/puppy/300",
			payload:     nil,
			HTTPCode:    http.StatusNotFound,
			expectedMsg: "Not Found: Apologies, for the puppy you seek does not exist\n"},
		{
			testName:    "Test POST Puppy",
			httpMethod:  "POST",
			url:         "/api/puppy/",
			payload:     []byte(`{"Breed": "Arcanine", "Colour": "Brown", "Value": 7000}`),
			HTTPCode:    http.StatusCreated,
			expectedMsg: "{\"id\":2,\"breed\":\"Arcanine\",\"colour\":\"Brown\",\"value\":7000}\n"},
		{
			testName:    "Test POST with negative value",
			httpMethod:  "POST",
			url:         "/api/puppy/",
			payload:     []byte(`{"Breed": "Arcanine", "Colour": "Brown", "Value": -7000}`),
			HTTPCode:    http.StatusUnprocessableEntity,
			expectedMsg: "PuppyStoreError 400: Sorry puppy value cannot be negative. The dog has to be worth something :)\n"},
		{
			testName:    "Test POST with JSON decode error",
			httpMethod:  "POST",
			url:         "/api/puppy/",
			payload:     []byte(`{"What: "I don't know", "When": "Now", "Where": "Somewhere"}`),
			HTTPCode:    http.StatusBadRequest,
			expectedMsg: "Bad Request: Invalid input, ensure ID and JSON are valid\n"},
		{
			testName:    "Test PUT puppy by ID",
			httpMethod:  "PUT",
			url:         "/api/puppy/1",
			payload:     []byte(`{"id": 1, "Breed": "Vulpix", "Colour": "Brown-ish", "Value": 2}`),
			HTTPCode:    http.StatusCreated,
			expectedMsg: "{\"id\":1,\"breed\":\"Vulpix\",\"colour\":\"Brown-ish\",\"value\":2}\n"},
		{
			testName:    "Test PUT puppy with invalid payload",
			httpMethod:  "PUT",
			url:         "/api/puppy/1",
			payload:     []byte(`{INVALID JSON}`),
			HTTPCode:    http.StatusUnprocessableEntity,
			expectedMsg: "Unprocessable Entity: Invalid input, ensure ID and JSON are valid\n"},
		{
			testName:    "Test PUT puppy with out of range id",
			httpMethod:  "PUT",
			url:         "/api/puppy/300",
			payload:     []byte(`{"id": 300, "Breed": "Pikachu", "Colour": "Yellow", "Value": 2100}`),
			HTTPCode:    http.StatusNotFound,
			expectedMsg: "Not Found: Apologies, for the puppy you seek does not exist\n"},
		{
			testName:    "Test PUT with invalid id",
			httpMethod:  "PUT",
			url:         "/api/puppy/foo",
			payload:     []byte(`{"Breed": "Liverbird", "Color": "Red", "Value": 20000}`),
			HTTPCode:    http.StatusBadRequest,
			expectedMsg: "Unprocessable Entity: Invalid input, ensure ID and JSON are valid\n"},
		{
			testName:    "Test PUT with non-int puppy value",
			httpMethod:  "PUT",
			url:         "/api/puppy/1",
			payload:     []byte(`{"Breed": "T-Rex", "Colour": "Green-ish", "Value": "blah"}`),
			HTTPCode:    http.StatusUnprocessableEntity,
			expectedMsg: "Unprocessable Entity: Invalid input, ensure ID and JSON are valid\n"},
		{
			testName:    "Test PUT with negative puppy value",
			httpMethod:  "PUT",
			url:         "/api/puppy/1",
			payload:     []byte(`{"Breed": "Sphinx", "Color": "Golden", "Value": -500000}`),
			HTTPCode:    http.StatusUnprocessableEntity,
			expectedMsg: "Unprocessable Entity: Invalid input, ensure ID and JSON are valid\n"},
		{
			testName:    "Test DELETE puppy by ID",
			httpMethod:  "DELETE",
			url:         "/api/puppy/1",
			payload:     nil,
			HTTPCode:    http.StatusOK,
			expectedMsg: "\"Puppy successfully deleted\"\n"},
		{
			testName:    "Test DELETE puppy with non-existent id",
			httpMethod:  "DELETE",
			url:         "/api/puppy/300",
			payload:     nil,
			HTTPCode:    http.StatusNotFound,
			expectedMsg: "Not Found: Apologies, for the puppy you seek does not exist\n"},
		{
			testName:    "Test DELETE puppy with invalid id",
			httpMethod:  "DELETE",
			url:         "/api/puppy/invalid",
			payload:     nil,
			HTTPCode:    http.StatusBadRequest,
			expectedMsg: "Bad Request: Invalid input, ensure ID and JSON are valid\n"},
	}
	for _, tc := range testCases {
		tc := tc // prevent scopelint error
		t.Run(tc.testName, func(t *testing.T) {
			runTestCase(t, tc.httpMethod, tc.url, tc.payload, tc.expectedMsg, tc.HTTPCode)
		})
	}
}
