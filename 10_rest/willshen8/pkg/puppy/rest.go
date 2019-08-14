package puppy

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// RestHandler provides a puppy storer to store the puppies
type RestHandler struct {
	// sync.Mutex
	store Storer
}

// NewRestHandler is a factory method that makes new storer
func NewRestHandler(storer Storer) *RestHandler {
	return &RestHandler{
		store: storer,
	}
}

// HandleGet gets the puppy by id and displays the results.
func (rh *RestHandler) HandleGet(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	puppy, readErr := rh.store.ReadPuppy(uint32(id))
	if readErr != nil {
		render.JSON(w, r, readErr)
	} else {
		render.JSON(w, r, puppy)
	}
}

// HandlePost post a new puppy and then display the results.
func (rh *RestHandler) HandlePost(w http.ResponseWriter, r *http.Request) {
	var p Puppy
	if decodeErr := render.DecodeJSON(r.Body, &p); decodeErr != nil {
		render.JSON(w, r, decodeErr)
	}
	_, createErr := rh.store.CreatePuppy(&p)
	if createErr != nil {
		render.JSON(w, r, createErr)
	} else {
		render.JSON(w, r, p)
	}
}

// HandlePut updates the existing puppy with new fields in the payload.
func (rh *RestHandler) HandlePut(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var p Puppy
	result, updateErr := rh.store.UpdatePuppy(uint32(id), &p)
	if updateErr != nil {
		render.JSON(w, r, updateErr)
	} else {
		render.JSON(w, r, result)
	}
}

// HandleDelete deletes the puppy with by id and display the results.
func (rh *RestHandler) HandleDelete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	result, deleteErr := rh.store.DeletePuppy(uint32(id))
	if deleteErr != nil {
		render.JSON(w, r, deleteErr)
	} else {
		render.JSON(w, r, result)
	}
}

// SetupRoutes provides the routes for this REST API.
func SetupRoutes(r chi.Router, rh RestHandler) {
	r.Get("/api/puppy/{id}", rh.HandleGet)
	r.Post("/api/puppy/", rh.HandlePost)
	r.Put("/api/puppy/{id}", rh.HandlePut)
	r.Delete("/api/puppy/{id}", rh.HandleDelete)
}
