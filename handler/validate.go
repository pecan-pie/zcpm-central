package handler

import "net/http"

func NewValidateHandler() http.Handler {
	return http.HandlerFunc(validate)
}

func validate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("validate function is not implemented"))
}
