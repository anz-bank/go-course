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
	store Storer
}

// NewRestHandler is a factory method that makes new storer
func NewRestHandler(storer Storer) *RestHandler {
	return &RestHandler{store: storer}
}

// HandleGet gets the puppy by id and displays the results.
func (rh *RestHandler) HandleGet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(http.StatusBadRequest)+": "+ErrInvalidInput.String(), http.StatusBadRequest)
		return
	}
	puppy, readErr := rh.store.ReadPuppy(uint32(id))
	if readErr != nil {
		w.WriteHeader(http.StatusNotFound)
		http.Error(w, http.StatusText(http.StatusNotFound)+": "+ErrNotFound.String(), http.StatusNotFound)
		return
	}
	render.JSON(w, r, puppy)
}

// HandlePost post a new puppy and then display the results.
func (rh *RestHandler) HandlePost(w http.ResponseWriter, r *http.Request) {
	var p Puppy
	if decodeErr := render.DecodeJSON(r.Body, &p); decodeErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(http.StatusBadRequest)+": "+ErrInvalidInput.String(), http.StatusBadRequest)
		return
	}
	_, createErr := rh.store.CreatePuppy(&p)
	if createErr != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest)+": "+ErrInvalidInput.String(), http.StatusBadRequest)
		return
	}
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, p)
}

// HandlePut updates the existing puppy with new fields in the payload.
func (rh *RestHandler) HandlePut(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(http.StatusBadRequest)+": "+ErrInvalidInput.String(), http.StatusBadRequest)
		return
	}
	var p Puppy
	if decodeErr := render.DecodeJSON(r.Body, &p); decodeErr != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity)+": "+ErrInvalidInput.String(),
			http.StatusUnprocessableEntity)
		return
	}

	updateErr := rh.store.UpdatePuppy(uint32(id), &p)
	if updateErr != nil {
		// if updateError is due to invalid id
		if updateErr.(*Error).Code == ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
			http.Error(w, http.StatusText(http.StatusNotFound)+": "+ErrNotFound.String(),
				http.StatusNotFound)
			return
		}
		// if updateError is due to invalid input
		if updateErr.(*Error).Code == ErrInvalidInput {
			w.WriteHeader(http.StatusBadRequest)
			http.Error(w, http.StatusText(http.StatusBadRequest)+": "+ErrInvalidInput.String(),
				http.StatusBadRequest)
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
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(http.StatusBadRequest)+": "+ErrInvalidInput.String(), http.StatusBadRequest)
		return
	}
	deleteErr := rh.store.DeletePuppy(uint32(id))
	if deleteErr != nil {
		w.WriteHeader(http.StatusNotFound)
		http.Error(w, http.StatusText(http.StatusNotFound)+": "+ErrInvalidInput.String(), http.StatusNotFound)
		return
	}
	render.JSON(w, r, "Puppy is deleted successfully")
	go notifyLostPuppy(id, w)
}

// notifyLostPuppy sends a POST response to puppy server when a puppy is deleted
func notifyLostPuppy(id int, w http.ResponseWriter) {
	payload := []byte(fmt.Sprintf(`{"ID":%d}`, id))
	resp, err := http.Post(lostPuppyServer, "application/json", bytes.NewReader(payload))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		logrus.Error("Something went wrong, the request was not successful")
		return
	}
	defer resp.Body.Close()
	logrus.Infof("Lost puppy with id of %d has a status code of %d", id, resp.StatusCode)
}

// SetupRoutes provides the routes for this REST API.
func SetupRoutes(r chi.Router, rh RestHandler) {
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
