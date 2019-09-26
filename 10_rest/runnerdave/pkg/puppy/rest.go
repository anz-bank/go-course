package puppy

import (
	"encoding/json"
	"net/http"
	"strconv"

	chi "github.com/go-chi/chi"
)

// RestServer store for puppies
type RestServer struct {
	Db Storer
}

// SetupRoutes setups the routes for the server
func (rs *RestServer) SetupRoutes(db Storer) *chi.Mux {
	r := chi.NewRouter()
	rs.Db = db

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
	if err == nil {
		p, rerr := rs.Db.ReadPuppy(int16(id))
		if rerr != nil {
			w.WriteHeader(http.StatusNotFound)
		}
		w.Header().Set("Content-Type", "application/json")
		bytes, _ := json.Marshal(p)
		i, _ := w.Write(bytes)
		w.Header().Set("Content-Length", string(i))
	}
}

func (rs *RestServer) createPuppy(w http.ResponseWriter, r *http.Request) {
	post := Puppy{}
	derr := json.NewDecoder(r.Body).Decode(&post)
	if derr != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	cerr := rs.Db.CreatePuppy(post)
	if cerr == nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (rs *RestServer) updatePuppy(w http.ResponseWriter, r *http.Request) {
	p := Puppy{}
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err == nil {
		derr := json.NewDecoder(r.Body).Decode(&p)
		if derr != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		uerr := rs.Db.UpdatePuppy(int16(id), &p)
		if uerr == nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func (rs *RestServer) deletePuppy(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err == nil {
		derr := rs.Db.DeletePuppy(int16(id))
		if derr != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
