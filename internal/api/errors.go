package api

import "net/http"

func WriteError(w http.ResponseWriter, err error) {
	// TODO: other error types such as user input
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("error"))
}
