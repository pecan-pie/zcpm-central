package handler

import (
	"net/http"
)

type RandomBody struct {
	from int
	to   int
}

type RandomResponse struct {
	number int
}

func NewRandomHandler() http.Handler {
	return http.HandlerFunc(random)
}

// gibt eine zufaellige Zahl in einem Wertebereich zurueck
func random(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("random function is not implemented"))
}
