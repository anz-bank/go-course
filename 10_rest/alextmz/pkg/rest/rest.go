package rest

import (
	"net/http"
	"path"
	"strconv"

	"github.com/anz-bank/go-course/10_rest/alextmz/pkg/puppy"
	"github.com/go-chi/render"
)

type HTTPHandler struct {
	Store puppy.Storer
}

// despite the name, http.Error writes any http code/string.
var writeHTTP = http.Error

func (ht HTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ht.handleGet(w, r)
	case "PUT":
		ht.handlePut(w, r)
	case "POST":
		ht.handlePost(w, r)
	case "DELETE":
		ht.handleDelete(w, r)
	}
}

func (ht *HTTPHandler) handleGet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(path.Base(r.RequestURI))
	if writeHTTPOnError(w, err) {
		return
	}

	puppy, err := ht.Store.ReadPuppy(id)

	if !writeHTTPOnError(w, err) {
		render.JSON(w, r, puppy)
	}
}

func (ht *HTTPHandler) handlePost(w http.ResponseWriter, r *http.Request) {
	var puppy puppy.Puppy

	if err := render.DecodeJSON(r.Body, &puppy); writeHTTPOnError(w, err) {
		return
	}

	err := ht.Store.CreatePuppy(&puppy)
	if !writeHTTPOnError(w, err) {
		w.WriteHeader(http.StatusCreated)
		render.JSON(w, r, puppy)
	}
}

func (ht *HTTPHandler) handlePut(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(path.Base(r.RequestURI))
	if writeHTTPOnError(w, err) {
		return
	}

	var puppy puppy.Puppy

	if err = render.DecodeJSON(r.Body, &puppy); writeHTTPOnError(w, err) {
		return
	}

	puppy.ID = id
	if err = ht.Store.UpdatePuppy(puppy); !writeHTTPOnError(w, err) {
		c := http.StatusOK
		writeHTTP(w, strconv.Itoa(c)+" "+http.StatusText(c), c)
	}
}

func (ht *HTTPHandler) handleDelete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(path.Base(r.RequestURI))
	if writeHTTPOnError(w, err) {
		return
	}

	err = ht.Store.DeletePuppy(id)

	if !writeHTTPOnError(w, err) {
		c := http.StatusOK
		writeHTTP(w, strconv.Itoa(c)+" "+http.StatusText(c), c)
	}
}

// writeHTTPOnError is a helper function - if there is no error, returns false,
// otherwise writes the error out and returns true.
func writeHTTPOnError(w http.ResponseWriter, err error) bool {
	if err != nil {
		var e int

		var m string

		if v, ok := err.(puppy.Error); ok {
			m = v.Error()
			// translates "puppy" errors to "http" errors
			switch v.Code {
			case puppy.ErrPuppyNotFoundOnRead:
				e = http.StatusNotFound
			case puppy.ErrPuppyNotFoundOnUpdate:
				e = http.StatusNotFound
			case puppy.ErrPuppyNotFoundOnDelete:
				e = http.StatusNotFound
			default:
				e = http.StatusBadRequest
			}
		} else {
			e = http.StatusBadRequest
			m = err.Error()
		}

		writeHTTP(w, strconv.Itoa(e)+" "+http.StatusText(e)+" : "+m, e)

		return true
	}

	return false
}
