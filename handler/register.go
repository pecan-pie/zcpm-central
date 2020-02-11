package handler

import "net/http"

func NewRegisterHandler() http.Handler {
	return http.HandlerFunc(register)
}

func register(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("register function is not implemented"))
}
