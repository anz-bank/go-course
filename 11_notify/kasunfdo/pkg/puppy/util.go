package puppy

import "net/http"

func sendErrorResponse(w http.ResponseWriter, statusCode int, err error) {
	w.WriteHeader(statusCode)
	http.Error(w, err.Error(), statusCode)
}
