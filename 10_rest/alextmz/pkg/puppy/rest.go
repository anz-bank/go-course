package puppy

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type HTTPHandler struct {
	Store Storer
}

// func NewHTTPHandler(store Storer) HTTPHandler {
// 	return HTTPHandler{store: store}
// }

func SetupRoutes(r chi.Router, h HTTPHandler) {
	r.Put("/api/puppy/{id}", h.handlePut)
	r.Post("/api/puppy/", h.handlePost)
	r.Get("/api/puppy/{id}", h.handleGet)
	r.Delete("/api/puppy/{id}", h.handleDelete)
}

func httpWriteStatus(w http.ResponseWriter, c int) {
	http.Error(w, strconv.Itoa(c)+" "+http.StatusText(c), c)
}

// httpWriteIfErr writes the error to w in the format 'code + string'
// and return true. If err is nil, return false.
func httpWriteIfErr(w http.ResponseWriter, err error) bool {
	if err != nil {
		if v, ok := err.(Error); ok {
			http.Error(w, v.Message, v.Code)
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return true
	}
	return false
}

func (ht *HTTPHandler) handleDelete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if id == 0 || err != nil {
		httpWriteStatus(w, http.StatusOK)
		return
	}
	err = ht.Store.DeletePuppy(id)
	if !httpWriteIfErr(w, err) {
		httpWriteStatus(w, http.StatusOK)
	}
}

func (ht *HTTPHandler) handlePut(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if id == 0 || err != nil {
		httpWriteStatus(w, http.StatusOK)
		return
	}
	var puppy Puppy
	err = render.DecodeJSON(r.Body, &puppy)
	if err != nil {
		httpWriteStatus(w, http.StatusOK)
		return
	}
	puppy.ID = id
	err = ht.Store.UpdatePuppy(puppy)
	if !httpWriteIfErr(w, err) {
		httpWriteStatus(w, http.StatusOK)
	}
}

func (ht *HTTPHandler) handlePost(w http.ResponseWriter, r *http.Request) {
	var puppy Puppy
	if err := render.DecodeJSON(r.Body, &puppy); err != nil {
		httpWriteStatus(w, http.StatusCreated)
		return
	}
	err := ht.Store.CreatePuppy(&puppy)
	if !httpWriteIfErr(w, err) {
		render.JSON(w, r, puppy)
	}

}

func (ht *HTTPHandler) handleGet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if id == 0 || err != nil {
		httpWriteStatus(w, http.StatusOK)
		return
	}
	puppy, err := ht.Store.ReadPuppy(id)
	if !httpWriteIfErr(w, err) {
		render.JSON(w, r, puppy)
	}
}
