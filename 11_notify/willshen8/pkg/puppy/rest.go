package puppy

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
)

const lostPuppyServer = "http://localhost:8888/api/lostpuppy/"

// RestHandler provides a puppy storer to store the puppies
type RestHandler struct {
	store  Storer
	client string
}

// NewRestHandler is a factory method that makes new storer
func NewRestHandler(storer Storer, lostpuppyServer string) *RestHandler {
	return &RestHandler{store: storer, client: lostpuppyServer}
}

// httpError is a helper function that output http status and error messages
func httpError(w http.ResponseWriter, httpStatus int, err ErrCode) {
	w.WriteHeader(httpStatus)
	http.Error(w, fmt.Sprintf("%s: %s", http.StatusText(httpStatus), err), httpStatus)
}

// HandleGet gets the puppy by id and displays the results.
func (rh *RestHandler) HandleGet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		httpError(w, http.StatusBadRequest, ErrInvalidInput)
		return
	}
	puppy, err := rh.store.ReadPuppy(uint32(id))
	if err != nil {
		httpError(w, http.StatusNotFound, ErrNotFound)
		return
	}
	render.JSON(w, r, puppy)
}

// HandlePost post a new puppy and then display the results.
func (rh *RestHandler) HandlePost(w http.ResponseWriter, r *http.Request) {
	var p Puppy
	if err := render.DecodeJSON(r.Body, &p); err != nil {
		httpError(w, http.StatusBadRequest, ErrInvalidInput)
		return
	}
	_, err := rh.store.CreatePuppy(&p)
	if err != nil {
		httpError(w, http.StatusBadRequest, ErrInvalidInput)
		return
	}
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, p)
}

// HandlePut updates the existing puppy with new fields in the payload.
func (rh *RestHandler) HandlePut(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		httpError(w, http.StatusBadRequest, ErrInvalidInput)
		return
	}
	var p Puppy
	if err := render.DecodeJSON(r.Body, &p); err != nil {
		httpError(w, http.StatusUnprocessableEntity, ErrInvalidInput)
		return
	}

	err = rh.store.UpdatePuppy(uint32(id), &p)
	if err != nil {
		if err.(*Error).Code == ErrNotFound {
			httpError(w, http.StatusNotFound, ErrNotFound)
			return
		}
		if err.(*Error).Code == ErrInvalidInput {
			httpError(w, http.StatusBadRequest, ErrInvalidInput)
			return
		}
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, "Puppy is updated successfully")
}

// HandleDelete deletes the puppy with by id and display the results.
func (rh *RestHandler) HandleDelete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		httpError(w, http.StatusBadRequest, ErrInvalidInput)
		return
	}
	err = rh.store.DeletePuppy(uint32(id))
	if err != nil {
		httpError(w, http.StatusNotFound, ErrInvalidInput)
		return
	}
	go rh.notifyLostPuppy(id)
	render.JSON(w, r, "Puppy is deleted successfully")
}

// notifyLostPuppy sends a POST response to puppy server when a puppy is deleted
func (rh *RestHandler) notifyLostPuppy(id int) {
	payload := []byte(fmt.Sprintf(`{"ID":%d}`, id))
	resp, err := http.Post(rh.client, "application/json", bytes.NewReader(payload))
	if err != nil {
		logrus.Error(err)
		return
	}
	defer resp.Body.Close()
	logrus.Infof("Lost puppy with id of %d has a status code of %d", id, resp.StatusCode)
}

// SetupRoutes provides the routes for this REST API.
func (rh *RestHandler) SetupRoutes(r chi.Router) {
	r.Get("/api/puppy/{id}", rh.HandleGet)
	r.Post("/api/puppy/", rh.HandlePost)
	r.Put("/api/puppy/{id}", rh.HandlePut)
	r.Delete("/api/puppy/{id}", rh.HandleDelete)
}

// SetupRouter takes the port number parsed from cmd and starts the server.
func SetupRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	return r
}
