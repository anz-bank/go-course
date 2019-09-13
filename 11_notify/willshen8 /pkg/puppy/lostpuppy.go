package puppy

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// Payload struct is the format that lost puppy server expects to see
type Payload struct {
	ID int `json:"id"`
}

// HandleLostPuppy processes the payload and determine the response based on id field
func HandleLostPuppy(w http.ResponseWriter, r *http.Request) {
	time.Sleep(2 * time.Second)
	var jsonData Payload
	if decodeErr := render.DecodeJSON(r.Body, &jsonData); decodeErr != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}
	if jsonData.ID%2 == 0 {
		w.WriteHeader(http.StatusCreated)
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
}

// SetupLostPuppyRoutes set up the routes for lost puppy server
func SetupLostPuppyRoutes(r chi.Router) {
	r.Post("/api/lostpuppy/", HandleLostPuppy)
}
