package puppy_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/anz-bank/go-course/11_notify/kasunfdo/pkg/puppy"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLostPuppyAPI(t *testing.T) {
	tests := []struct {
		name       string
		httpMethod string
		url        string
		payload    []byte
		httpCode   int
		expected   string
	}{
		{
			name:       "POST even id",
			httpMethod: "POST",
			url:        "/api/lostpuppy/",
			payload:    []byte(`{"id": 2}`),
			httpCode:   http.StatusCreated,
			expected:   "{\"Status\":201}\n",
		},
		{
			name:       "POST odd id",
			httpMethod: "POST",
			url:        "/api/lostpuppy/",
			payload:    []byte(`{"id": 1}`),
			httpCode:   http.StatusInternalServerError,
			expected:   "{\"Status\":500}\n",
		},
		{
			name:       "POST invalid request",
			httpMethod: "POST",
			url:        "/api/lostpuppy/",
			payload:    []byte(`{"id": "foo"}`),
			httpCode:   http.StatusBadRequest,
			expected: "invalid input: \n\tjson: " +
				"cannot unmarshal string into Go struct field LostPuppyRequest.id of type uint64\n",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			request, err := http.NewRequest(test.httpMethod, test.url, bytes.NewBuffer(test.payload))
			require.NoError(t, err)

			router := chi.NewRouter()
			puppy.NewLostAPIHandler().WireRoutes(router)
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
