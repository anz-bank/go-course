package puppy

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// APIHandler implements REST API handlers of Puppy.
type APIHandler struct {
	store Storer
}

// NewAPIHandler creates a APIHandler with given storer.
func NewAPIHandler(storer Storer) *APIHandler {
	return &APIHandler{store: storer}
}

// HandleGetPuppyByID retrieves puppy by id.
func (a *APIHandler) HandleGetPuppyByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		sendErrorResponse(w, http.StatusBadRequest, ErrorEf(ErrInvalid, err, ""))
		return
	}

	puppy, err := a.store.ReadPuppy(id)
	if err != nil {
		sendErrorResponse(w, http.StatusNotFound, err)
		return
	}
	render.JSON(w, r, puppy)
}

// HandlePostPuppy adds puppy into the store.
func (a *APIHandler) HandlePostPuppy(w http.ResponseWriter, r *http.Request) {
	var p Puppy
	if err := render.DecodeJSON(r.Body, &p); err != nil {
		sendErrorResponse(w, http.StatusBadRequest, ErrorEf(ErrInvalid, err, ""))
		return
	}

	id, err := a.store.CreatePuppy(p)
	if err != nil {
		sendErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	p.ID = id
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, p)
}

// HandlePutPuppy updates puppy in the store.
func (a *APIHandler) HandlePutPuppy(w http.ResponseWriter, r *http.Request) {
	var p Puppy
	if err := render.DecodeJSON(r.Body, &p); err != nil {
		sendErrorResponse(w, http.StatusBadRequest, ErrorEf(ErrInvalid, err, ""))
		return
	}

	if err := a.store.UpdatePuppy(p); err != nil {
		sendErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, "puppy updated")
}

// HandleDeletePuppy deletes puppy in store by id.
func (a *APIHandler) HandleDeletePuppy(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		sendErrorResponse(w, http.StatusBadRequest, ErrorEf(ErrInvalid, err, ""))
		return
	}
	if err := a.store.DeletePuppy(id); err != nil {
		sendErrorResponse(w, http.StatusNotFound, err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, "puppy deleted")
}

// WireRoutes route requests to corresponding REST API handler method.
func (a *APIHandler) WireRoutes(r chi.Router) {
	r.Get("/api/puppy/{id}", a.HandleGetPuppyByID)
	r.Post("/api/puppy/", a.HandlePostPuppy)
	r.Put("/api/puppy/", a.HandlePutPuppy)
	r.Delete("/api/puppy/{id}", a.HandleDeletePuppy)
}

func sendErrorResponse(w http.ResponseWriter, statusCode int, err error) {
	w.WriteHeader(statusCode)
	http.Error(w, err.Error(), statusCode)
}
