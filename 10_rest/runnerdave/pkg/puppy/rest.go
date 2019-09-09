package puppy

import (
	"encoding/json"
	"net/http"
	"strconv"

	chi "github.com/go-chi/chi"
)

// RestStorer store for puppies
type RestStorer struct {
	Db Storer
}

const (
	getPuppyRoute    = "/api/puppy/{id}"
	postPuppyRoute   = "/api/puppy/"
	putPuppyRoute    = "/api/puppy/{id}"
	deletePuppyRoute = "/api/puppy/{id}"
)

// GetPuppyRoute Route for the GET method
func (rs *RestStorer) GetPuppyRoute() string {
	return getPuppyRoute
}

// PostPuppyRoute Route for the POST method
func (rs *RestStorer) PostPuppyRoute() string {
	return postPuppyRoute
}

// PutPuppyRoute Route for the PUT method
func (rs *RestStorer) PutPuppyRoute() string {
	return putPuppyRoute
}

// DeletePuppyRoute Route for the PUT method
func (rs *RestStorer) DeletePuppyRoute() string {
	return deletePuppyRoute
}

// GetPuppy HandlerFunction
func (rs *RestStorer) GetPuppy(w http.ResponseWriter, r *http.Request) {
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

// CreatePuppy Handler function to create puppies
func (rs *RestStorer) CreatePuppy(w http.ResponseWriter, r *http.Request) {
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

// UpdatePuppy Updates an existing Puppy
func (rs *RestStorer) UpdatePuppy(w http.ResponseWriter, r *http.Request) {
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

// DeletePuppy Deletes an existing Puppy from the database
func (rs *RestStorer) DeletePuppy(w http.ResponseWriter, r *http.Request) {
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
