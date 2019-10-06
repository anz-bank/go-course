package puppy_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/anz-bank/go-course/11_notify/kasunfdo/pkg/puppy"
	"github.com/anz-bank/go-course/11_notify/kasunfdo/pkg/puppy/store"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func getTestRouter(t *testing.T) *chi.Mux {
	storer := store.NewMapStore()
	_, err := storer.CreatePuppy(puppy.Puppy{ID: 1, Breed: "Labrador", Colour: "Cream", Value: 2000})
	require.NoError(t, err)

	router := chi.NewRouter()
	puppy.NewAPIHandler(storer, "").WireRoutes(router)

	return router
}

func TestAPI(t *testing.T) {
	tests := []struct {
		name       string
		httpMethod string
		url        string
		payload    []byte
		httpCode   int
		expected   string
	}{
		{
			name:       "GET puppy ID valid",
			httpMethod: "GET",
			url:        "/api/puppy/1",
			payload:    nil,
			httpCode:   http.StatusOK,
			expected:   `{"id":1,"breed":"Labrador","colour":"Cream","value":2000}` + "\n"},
		{
			name:       "GET puppy ID invalid id",
			httpMethod: "GET",
			url:        "/api/puppy/0",
			payload:    nil,
			httpCode:   http.StatusNotFound,
			expected:   "not found: puppy with id: 0 is not found\n"},
		{
			name:       "GET puppy ID invalid input",
			httpMethod: "GET",
			url:        "/api/puppy/foo",
			payload:    nil,
			httpCode:   http.StatusBadRequest,
			expected:   "invalid input: \n\tstrconv.ParseUint: parsing \"foo\": invalid syntax\n"},
		{
			name:       "POST puppy success",
			httpMethod: "POST",
			url:        "/api/puppy/",
			payload:    []byte(`{"breed": "Rottweiler", "colour": "Brown", "value": 2100}`),
			httpCode:   http.StatusCreated,
			expected:   `{"id":2,"breed":"Rottweiler","colour":"Brown","value":2100}` + "\n"},
		{
			name:       "POST puppy negative val",
			httpMethod: "POST",
			url:        "/api/puppy/",
			payload:    []byte(`{"breed": "Rottweiler", "colour": "Brown", "value": -2100}`),
			httpCode:   http.StatusBadRequest,
			expected:   "invalid input: value of puppy is negative\n"},
		{
			name:       "POST puppy invalid data",
			httpMethod: "POST",
			url:        "/api/puppy/",
			payload:    []byte(`Gibberish`),
			httpCode:   http.StatusBadRequest,
			expected:   "invalid input: \n\tinvalid character 'G' looking for beginning of value\n"},
		{
			name:       "PUT puppy success",
			httpMethod: "PUT",
			url:        "/api/puppy/",
			payload:    []byte(`{"id":1,"breed": "Rottweiler", "colour": "Brown", "value": 2100}`),
			httpCode:   http.StatusOK,
			expected:   "{\"Status\":200,\"Msg\":\"puppy updated\"}\n"},
		{
			name:       "PUT puppy negative val",
			httpMethod: "PUT",
			url:        "/api/puppy/",
			payload:    []byte(`{"id":1,"breed": "Rottweiler", "colour": "Brown", "value": -2100}`),
			httpCode:   http.StatusBadRequest,
			expected:   "invalid input: value of puppy is negative\n"},
		{
			name:       "PUT puppy invalid id",
			httpMethod: "PUT",
			url:        "/api/puppy/",
			payload:    []byte(`{"id":0,"breed": "Rottweiler", "colour": "Brown", "value": 2100}`),
			httpCode:   http.StatusBadRequest,
			expected:   "not found: puppy with id: 0 is not found\n"},
		{
			name:       "PUT puppy invalid data",
			httpMethod: "PUT",
			url:        "/api/puppy/",
			payload:    []byte(`Gibberish`),
			httpCode:   http.StatusBadRequest,
			expected:   "invalid input: \n\tinvalid character 'G' looking for beginning of value\n"},
		{
			name:       "DELETE puppy success",
			httpMethod: "DELETE",
			url:        "/api/puppy/1",
			payload:    nil,
			httpCode:   http.StatusOK,
			expected:   "{\"Status\":200,\"Msg\":\"puppy deleted\"}\n"},
		{
			name:       "DELETE puppy invalid id",
			httpMethod: "DELETE",
			url:        "/api/puppy/0",
			payload:    nil,
			httpCode:   http.StatusNotFound,
			expected:   "not found: puppy with id: 0 is not found\n"},
		{
			name:       "DELETE puppy invalid input",
			httpMethod: "DELETE",
			url:        "/api/puppy/foo",
			payload:    nil,
			httpCode:   http.StatusBadRequest,
			expected:   "invalid input: \n\tstrconv.ParseUint: parsing \"foo\": invalid syntax\n"},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			request, err := http.NewRequest(test.httpMethod, test.url, bytes.NewBuffer(test.payload))
			require.NoError(t, err)

			router := getTestRouter(t)
			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder, request)
			response := recorder.Result()
			defer response.Body.Close()

			body, err := ioutil.ReadAll(response.Body)
			require.NoError(t, err)
			actual := string(body)

			assert.Equal(t, test.httpCode, response.StatusCode)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestReportLostPuppy(t *testing.T) {
	puppy.ReportLostPuppy(1, "http://httpbin.org/post")
}
