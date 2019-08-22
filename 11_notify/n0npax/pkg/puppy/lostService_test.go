package puppy

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	tassert "github.com/stretchr/testify/assert"
)

// LostPuppyBackend
func TestLostPuppy(t *testing.T) {
	SlowRequestDuration = time.Duration(0)
	t.Run("Odd", LostPupptOdd)
	t.Run("Even", LostPupptEven)
	t.Run("Err", LostPupptErr)
}

func LostPupptErr(t *testing.T) {
	assert := tassert.New(t)
	router := LostPuppyBackend()
	payload := string(`{"id": "llama"}`)
	w := postReq(router, "/api/lostpuppy/", strings.NewReader(payload))
	assert.Equal(400, w.Code)
}

func LostPupptEven(t *testing.T) {
	for i := -3; i <= 13; i++ {
		num := i * 2
		t.Run(fmt.Sprintf("even: %d", num), func(t *testing.T) {
			assert := tassert.New(t)
			router := LostPuppyBackend()
			payload := fmt.Sprintf(`{"id": %d}`, num)
			w := postReq(router, "/api/lostpuppy/", strings.NewReader(payload))
			assert.Equal(201, w.Code)
		})
	}
}

func LostPupptOdd(t *testing.T) {
	for i := -3; i <= 13; i++ {
		num := i*2 + 1
		t.Run(fmt.Sprintf("even: %d", num), func(t *testing.T) {
			assert := tassert.New(t)
			router := LostPuppyBackend()
			payload := fmt.Sprintf(`{"id": %d}`, num)
			w := postReq(router, "/api/lostpuppy/", strings.NewReader(payload))
			assert.Equal(500, w.Code)
		})
	}
}

func postReq(r http.Handler, path string, payload io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest("POST", path, payload)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
