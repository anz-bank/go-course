package puppy

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func runTest(t *testing.T, payload []byte, expectedBody string, expectedHTTPCode int) {
	req, err := http.NewRequest("POST", "/api/lostpuppy/", bytes.NewBuffer(payload))
	if err != nil {
		require.NoError(t, err)
	}

	// setup router
	router := SetupRouter()

	//setup routes
	SetupLostPuppyRoutes(router)
	// Setup recorder
	rr := httptest.NewRecorder()

	// Execute
	router.ServeHTTP(rr, req)

	// Get actual response
	resp := rr.Result()
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)
	bodyString := string(body)

	// check for http status
	assert.Equal(t, expectedHTTPCode, resp.StatusCode)
	// check for body
	assert.Equal(t, expectedBody, bodyString)
}

func TestLostPuppyAPI(t *testing.T) {
	tests := []struct {
		testName    string
		payload     []byte
		HTTPCode    int
		expectedMsg string
	}{
		{
			testName:    "Test POST with nil payload",
			payload:     nil,
			HTTPCode:    http.StatusUnprocessableEntity,
			expectedMsg: "Unprocessable Entity\n",
		},
		{
			testName:    "Test POST with odd ID",
			payload:     []byte(`{"id": 1}`),
			HTTPCode:    http.StatusInternalServerError,
			expectedMsg: "",
		},
		{
			testName:    "Test POST with even ID",
			payload:     []byte(`{"ID": 2}`),
			HTTPCode:    http.StatusCreated,
			expectedMsg: "",
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.testName, func(t *testing.T) {
			runTest(t, tc.payload, tc.expectedMsg, tc.HTTPCode)
		})
	}
}
