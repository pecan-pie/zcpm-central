package handler

import "net/http"

func NewValidateHandler() http.Handler {
	return http.HandlerFunc(validate)
}

type ValidateBody struct {
	hosts []string
}

type ValidateResponse struct {
	hosts []string
}

// validiert ob ein Host aus einer Liste von Hosts valide ist
// gibt eine Liste mit validen Hosts an den Aufrufer zurueck
// valide ist jeder Hosts, der registiert ist
func validate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("validate function is not implemented"))
}
