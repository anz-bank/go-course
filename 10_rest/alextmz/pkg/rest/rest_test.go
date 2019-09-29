package rest

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/anz-bank/go-course/10_rest/alextmz/pkg/puppy/store"
	"github.com/stretchr/testify/assert"
)

type HTTPHandlerTest struct {
	HTTPHandler
}

// func printpuppies(s puppy.Storer, n int) {
// 	for i := 1; i <= n; i++ {
// 		p, err := s.ReadPuppy(i)
// 		if err != nil {
// 			fmt.Printf("error printing puppies: %v\n", err)
// 			return
// 		}

// 		fmt.Printf("Printing puppy id %d: %#v\n", i, p)
// 	}
// }

func TestHandleGet(t *testing.T) {
	var tests = []struct {
		name, argmethod, argurl, argjson string
		wantcode                         int
		wantbody                         string
	}{
		{
			name:      "GET a not yet existing puppy 1",
			argmethod: "GET",
			argurl:    "1",
			wantcode:  http.StatusNotFound,
			wantbody:  "404 Not Found : puppy with ID 1 being read does not exist\n",
		},
		{
			name:      "FIXTURE: POST a puppy",
			argmethod: "POST",
			argjson:   `{"breed":"OKAPuppy","colour":"Rainbow","value":4}`,
			wantcode:  http.StatusCreated,
			wantbody:  `{"id":1,"breed":"OKAPuppy","colour":"Rainbow","value":4}` + "\n",
		},
		{
			name:      "FIXTURE: POST a puppy",
			argmethod: "POST",
			argjson:   `{"breed":"OKEPuppy","colour":"Invisible","value":2}`,
			wantcode:  http.StatusCreated,
			wantbody:  `{"id":2,"breed":"OKEPuppy","colour":"Invisible","value":2}` + "\n",
		},
		{
			name:      "GET valid, existing puppy 1",
			argmethod: "GET",
			argurl:    "1",
			wantcode:  http.StatusOK,
			wantbody:  `{"id":1,"breed":"OKAPuppy","colour":"Rainbow","value":4}` + "\n",
		},
		{
			name:      "GET valid, existing puppy 2",
			argmethod: "GET",
			argurl:    "2",
			wantcode:  http.StatusOK,
			wantbody:  `{"id":2,"breed":"OKEPuppy","colour":"Invisible","value":2}` + "\n",
		},
		{
			name:      "GET a non-existing puppy 42",
			argmethod: "GET",
			argurl:    "42",
			wantcode:  http.StatusNotFound,
			wantbody:  "404 Not Found : puppy with ID 42 being read does not exist\n",
		},
		{
			name:      "GET an invalid puppy -1",
			argmethod: "GET",
			argurl:    "-1",
			wantcode:  http.StatusNotFound,
			wantbody:  "404 Not Found : puppy with ID -1 being read does not exist\n",
		},
		{
			name:      "GET an invalid puppy 0",
			argmethod: "GET",
			argurl:    "0",
			wantcode:  http.StatusNotFound,
			wantbody:  "404 Not Found : puppy with ID 0 being read does not exist\n",
		},
		{
			name:      "GET an invalid puppy X",
			argmethod: "GET",
			argurl:    "X",
			wantcode:  http.StatusBadRequest,
			wantbody:  "400 Bad Request : strconv.Atoi: parsing \"X\": invalid syntax\n",
		},
	}

	var h HTTPHandlerTest

	for _, v := range tests {
		test := v
		t.Run(test.name, func(t *testing.T) {
			gotcode, gotbody, err := h.runRequest(test.argmethod, test.argurl, test.argjson)
			assert.Equal(t, test.wantcode, gotcode)
			assert.Equal(t, test.wantbody, gotbody)
			assert.NoError(t, err)
		})
	}
}

func TestHandlePost(t *testing.T) {
	var h HTTPHandlerTest

	var tests = []struct {
		name      string
		argmethod string
		argurl    string
		argjson   string
		wantcode  int
		wantbody  string
	}{
		{
			name:      "FIXTURE: POST 1st valid puppy",
			argmethod: "POST",
			argjson:   `{"breed":"OKPuppy","colour":"Rainbow","value":4}`,
			wantcode:  http.StatusCreated,
			wantbody:  `{"id":1,"breed":"OKPuppy","colour":"Rainbow","value":4}` + "\n",
		},
		{
			name:      "FIXTURE: POST 2nd valid puppy",
			argmethod: "POST",
			argjson:   `{"breed":"OKPuppy2","colour":"Invisible","value":2}`,
			wantcode:  http.StatusCreated,
			wantbody:  `{"id":2,"breed":"OKPuppy2","colour":"Invisible","value":2}` + "\n",
		},
		{
			name:      "POST invalid puppy JSON",
			argmethod: "POST",
			argjson:   `{"breed":"InvalidJsonPuppy",colour:Rainbow}`,
			wantcode:  http.StatusBadRequest,
			wantbody:  "400 Bad Request : invalid character 'c' looking for beginning of object key string\n",
		},
		{
			name:      "POST invalid already-identified puppy",
			argmethod: "POST",
			argjson:   `{"ID": 42, "breed":"OKPuppy","colour":"Rainbow","value":5}`,
			wantcode:  http.StatusBadRequest,
			wantbody:  "400 Bad Request : puppy already initialized with ID 42\n",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			gotcode, gotbody, err := h.runRequest(test.argmethod, test.argurl, test.argjson)
			assert.Equal(t, test.wantcode, gotcode)
			assert.Equal(t, test.wantbody, gotbody)
			assert.NoError(t, err)
		})
	}
}

func TestHandlePut(t *testing.T) {
	var tests = []struct {
		name      string
		argmethod string
		argurl    string
		argjson   string
		wantcode  int
		wantbody  string
	}{
		{
			name:      "FIXTURE: POST initial puppy 1",
			argmethod: "POST",
			argjson:   `{"breed":"OKPuppy","colour":"Rainbow","value":5}`,
			wantcode:  http.StatusCreated,
			wantbody:  `{"id":1,"breed":"OKPuppy","colour":"Rainbow","value":5}` + "\n",
		},
		{
			name:      "FIXTURE: POST valid puppy 2",
			argmethod: "POST",
			argjson:   `{"breed":"OKPuppy2","colour":"Invisible","value":5}`,
			wantcode:  http.StatusCreated,
			wantbody:  `{"id":2,"breed":"OKPuppy2","colour":"Invisible","value":5}` + "\n",
		},
		{
			name:      "GET the valid, existing puppy 1",
			argmethod: "GET",
			argurl:    "1",
			wantcode:  http.StatusOK,
			wantbody:  `{"id":1,"breed":"OKPuppy","colour":"Rainbow","value":5}` + "\n",
		},
		{
			name:      "PUT changed puppy 1",
			argmethod: "PUT",
			argurl:    "1",
			argjson:   `{"breed":"OKPuppy","colour":"HalfRainbow","value":20}`,
			wantcode:  http.StatusOK,
			wantbody:  "200 OK\n",
		},
		{
			name:      "GET the changed, valid, existing puppy 1",
			argmethod: "GET",
			argurl:    "1",
			wantcode:  http.StatusOK,
			wantbody:  `{"id":1,"breed":"OKPuppy","colour":"HalfRainbow","value":20}` + "\n",
		},
		{
			name:      "PUT changed puppy 2",
			argmethod: "PUT",
			argurl:    "2",
			argjson:   `{"breed":"Unknown","colour":"Blue","value":0}`,
			wantcode:  http.StatusOK,
			wantbody:  "200 OK\n",
		},
		{
			name:      "GET the changed, valid, existing puppy 2",
			argmethod: "GET",
			argurl:    "2",
			wantcode:  http.StatusOK,
			wantbody:  `{"id":2,"breed":"Unknown","colour":"Blue","value":0}` + "\n",
		},
		{
			name:      "PUT a changed, invalid puppy 2",
			argmethod: "PUT",
			argurl:    "2",
			argjson:   `{"breed":"Unknown","colour":"Blue","value":-1}`,
			wantcode:  http.StatusBadRequest,
			wantbody:  "400 Bad Request : trying to update a puppy with a negative value (-1.00)\n",
		},
		{
			name:      "GET make sure puppy 2 wasnt modified",
			argmethod: "GET",
			argurl:    "2",
			wantcode:  http.StatusOK,
			wantbody:  `{"id":2,"breed":"Unknown","colour":"Blue","value":0}` + "\n",
		},
		{
			name:      "PUT a changed, invalid puppy X",
			argmethod: "PUT",
			argurl:    "X",
			argjson:   `{"breed":"OKPuppy","colour":"HalfRainbow","value":20}`,
			wantcode:  http.StatusBadRequest,
			wantbody:  "400 Bad Request : strconv.Atoi: parsing \"X\": invalid syntax\n",
		},
		{
			name:      "PUT a changed, invalid JSON puppy 1",
			argmethod: "PUT",
			argurl:    "1",
			argjson:   `"breed":"OKPuppy","colour":"HalfRainbow","value":20}`,
			wantcode:  http.StatusBadRequest,
			wantbody:  "400 Bad Request : json: cannot unmarshal string into Go value of type puppy.Puppy\n",
		},
		{
			name: "GET make sure puppy 1 wasnt modified	",
			argmethod: "GET",
			argurl:    "1",
			wantcode:  http.StatusOK,
			wantbody:  `{"id":1,"breed":"OKPuppy","colour":"HalfRainbow","value":20}` + "\n",
		},
		{
			name:      "PUT non-existing puppy 99",
			argmethod: "PUT",
			argurl:    "99",
			argjson:   `{"breed":"Unknown","colour":"Blue","value":0}`,
			wantcode:  http.StatusNotFound,
			wantbody:  "404 Not Found : puppy with ID 99 being updated does not exist\n",
		},
	}

	var h HTTPHandlerTest

	for _, v := range tests {
		test := v
		t.Run(test.name, func(t *testing.T) {
			gotcode, gotbody, err := h.runRequest(test.argmethod, test.argurl, test.argjson)
			assert.Equal(t, test.wantcode, gotcode)
			assert.Equal(t, test.wantbody, gotbody)
			assert.NoError(t, err)
		})
	}
}

func TestHandleDelete(t *testing.T) {
	var tests = []struct {
		name, argmethod, argurl, argjson string
		wantcode                         int
		wantbody                         string
	}{
		{
			name:      "DELETE a puppy that never existed",
			argmethod: "DELETE",
			argurl:    "1",
			wantcode:  http.StatusNotFound,
			wantbody:  "404 Not Found : puppy with ID 1 being deleted does not exist\n",
		},
		{
			name:      "FIXTURE: POST initial puppy 1",
			argmethod: "POST",
			argjson:   `{"breed":"OKPuppy","colour":"Rainbow","value":5}`,
			wantcode:  http.StatusCreated,
			wantbody:  `{"id":1,"breed":"OKPuppy","colour":"Rainbow","value":5}` + "\n",
		},
		{
			name:      "FIXTURE: POST puppy 2",
			argmethod: "POST",
			argjson:   `{"breed":"OKPuppy2","colour":"Red","value":1}`,
			wantcode:  http.StatusCreated,
			wantbody:  `{"id":2,"breed":"OKPuppy2","colour":"Red","value":1}` + "\n",
		},
		{
			name:      "DELETE a valid puppy 1",
			argmethod: "DELETE",
			argurl:    "1",
			wantcode:  http.StatusOK,
			wantbody:  "200 OK\n",
		},
		{
			name:      "DELETE already deleted puppy 1",
			argmethod: "DELETE",
			argurl:    "1",
			wantcode:  http.StatusNotFound,
			wantbody:  "404 Not Found : puppy with ID 1 being deleted does not exist\n",
		},
		{
			name:      "FIXTURE: POST puppy 3",
			argmethod: "POST",
			argjson:   `{"breed":"OKPuppy","colour":"Rainbow","value":5}`,
			wantcode:  http.StatusCreated,
			wantbody:  `{"id":3,"breed":"OKPuppy","colour":"Rainbow","value":5}` + "\n",
		},
		{
			name:      "DELETE on invalid puppy 0",
			argmethod: "DELETE",
			argurl:    "0",
			wantcode:  http.StatusNotFound,
			wantbody:  "404 Not Found : puppy with ID 0 being deleted does not exist\n",
		},
		{
			name:      "DELETE on invalid puppy -1",
			argmethod: "DELETE",
			argurl:    "-1",
			wantcode:  http.StatusNotFound,
			wantbody:  "404 Not Found : puppy with ID -1 being deleted does not exist\n",
		},
		{
			name:      "DELETE on invalid puppy X",
			argmethod: "DELETE",
			argurl:    "X",
			wantcode:  http.StatusBadRequest,
			wantbody:  "400 Bad Request : strconv.Atoi: parsing \"X\": invalid syntax\n",
		},
		{
			name:      "DELETE a valid puppy 2",
			argmethod: "DELETE",
			argurl:    "2",
			wantcode:  http.StatusOK,
			wantbody:  "200 OK\n",
		},
	}

	var h HTTPHandlerTest

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			gotcode, gotbody, err := h.runRequest(test.argmethod, test.argurl, test.argjson)
			assert.Equal(t, test.wantcode, gotcode)
			assert.Equal(t, test.wantbody, gotbody)
			assert.NoError(t, err)
		})
	}
}

func (h *HTTPHandlerTest) runRequest(method, urlparam, puppystr string) (int, string, error) {
	url := "/api/puppy/" + urlparam

	if h.Store == nil {
		h.Store = store.NewMapStore()
	}

	puppyreader := strings.NewReader(puppystr)
	req := httptest.NewRequest(method, url, puppyreader)
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)

	resp := rr.Result()
	defer resp.Body.Close()

	bbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, "", errors.New("could not read response body, error " + err.Error())
	}

	return resp.StatusCode, string(bbody), nil
}
