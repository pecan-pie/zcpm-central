package handler

import (
	"net/http"
)

func NewRandomHandler() http.Handler {
	return http.HandlerFunc(random)
}

func random(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("random function is not implemented"))
}
