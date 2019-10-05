package puppy

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// LostPuppyRequest struct is the format that lost puppy server expects to see
type LostPuppyRequest struct {
	ID int `json:"id"`
}

// HandleLostPuppy processes the payload and determine the response based on id field
func HandleLostPuppy(w http.ResponseWriter, r *http.Request) {
	time.Sleep(2 * time.Second)
	var jsonData LostPuppyRequest
	if err := render.DecodeJSON(r.Body, &jsonData); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}
	if jsonData.ID%2 == 1 {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// SetupLostPuppyRoutes set up the routes for lost puppy server
func SetupLostPuppyRoutes(r chi.Router) {
	r.Post("/api/lostpuppy/", HandleLostPuppy)
}
