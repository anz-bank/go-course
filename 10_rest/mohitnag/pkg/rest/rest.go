package rest

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/anz-bank/go-course/10_rest/mohitnag/pkg/puppy"
	"github.com/anz-bank/go-course/10_rest/mohitnag/pkg/puppy/store"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

type Handler struct {
	storer store.Storer
	router chi.Router
}

func (rh *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rh.router.ServeHTTP(w, r)
}

func NewRestHandler(s store.Storer) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	rh := Handler{
		storer: s,
		router: r}
	rh.handleRoutes()
	return &rh
}

func (rh *Handler) handleRoutes() {
	r := rh.router
	r.Route("/api/puppy", func(r chi.Router) {
		r.Route("/", func(r chi.Router) {
			r.Post("/", rh.handlePostPuppy)
			r.Put("/", rh.handleUpdatePuppy)
		})
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", rh.handleGetPuppy)
			r.Delete("/", rh.handleDeletePuppy)
		})
	})
}

func (rh *Handler) handlePostPuppy(w http.ResponseWriter, r *http.Request) {
	var puppy puppy.Puppy
	if err := render.DecodeJSON(r.Body, &puppy); err != nil {
		renderDecodeJSONErrorResponse(w, r, err)
		return
	}
	if err := rh.storer.CreatePuppy(puppy); err != nil {
		renderErrorResponse(w, r, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (rh *Handler) handleGetPuppy(w http.ResponseWriter, r *http.Request) {
	var puppy puppy.Puppy
	ID := chi.URLParam(r, "id")
	puppyID, err := strconv.Atoi(ID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	puppy, err = rh.storer.ReadPuppy(uint32(puppyID))
	if err != nil {
		renderErrorResponse(w, r, err)
		return
	}
	render.JSON(w, r, puppy)
}

func (rh *Handler) handleUpdatePuppy(w http.ResponseWriter, r *http.Request) {
	var puppy puppy.Puppy
	if err := render.DecodeJSON(r.Body, &puppy); err != nil {
		renderDecodeJSONErrorResponse(w, r, err)
		return
	}
	err := rh.storer.UpdatePuppy(puppy)
	if err != nil {
		renderErrorResponse(w, r, err)
		return
	}
	render.JSON(w, r, puppy)
}

func (rh *Handler) handleDeletePuppy(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "id")
	puppyID, err := strconv.Atoi(ID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = rh.storer.DeletePuppy(uint32(puppyID))
	if err != nil {
		renderErrorResponse(w, r, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func renderErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	status := http.StatusInternalServerError
	body := map[string]interface{}{"error": err.Error()}
	if puppyError, ok := err.(*puppy.Error); ok {
		body["error"] = puppyError.Message
		body["code"] = puppyError.Code
		switch puppyError.Code {
		case puppy.NotFound:
			status = http.StatusNotFound
		case puppy.Invalid, puppy.Duplicate:
			status = http.StatusBadRequest
		}
	}
	render.Status(r, status)
	render.JSON(w, r, body)
}

func renderDecodeJSONErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		err = puppy.ErrorF(puppy.Invalid, "Invalid JSON for required data type")
	}
	renderErrorResponse(w, r, err)
}
