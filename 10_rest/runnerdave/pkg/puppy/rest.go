package puppy

import (
	"encoding/json"
	"net/http"
	"strconv"

	chi "github.com/go-chi/chi"
)

// RestServer store for puppies
type RestServer struct {
	DB Storer
}

// SetupRoutes setups the routes for the server
func (rs *RestServer) SetupRoutes(db Storer) *chi.Mux {
	r := chi.NewRouter()
	rs.DB = db

	r.Route("/api/puppy", func(r chi.Router) {
		r.Post("/", rs.createPuppy)
		r.Route("/{id}", func(r chi.Router) {
			r.Put("/", rs.updatePuppy)
			r.Delete("/", rs.deletePuppy)
			r.Get("/", rs.getPuppy)
		})
	})

	return r
}

func (rs *RestServer) getPuppy(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	p, err := rs.DB.ReadPuppy(int16(id))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	bytes, _ := json.Marshal(p)
	_, _ = w.Write(bytes)
}

func (rs *RestServer) createPuppy(w http.ResponseWriter, r *http.Request) {
	p := Puppy{}
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := rs.DB.CreatePuppy(p); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (rs *RestServer) updatePuppy(w http.ResponseWriter, r *http.Request) {
	p := Puppy{}
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err = json.NewDecoder(r.Body).Decode(&p); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err = rs.DB.UpdatePuppy(int16(id), &p); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (rs *RestServer) deletePuppy(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err = rs.DB.DeletePuppy(int16(id)); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
