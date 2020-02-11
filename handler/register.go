package handler

import "net/http"

func NewRegisterHandler() http.Handler {
	return http.HandlerFunc(register)
}

// registiert einen anfragenden Host
// sollte anhand der IP oder des Hostnamen des Hosts gemacht werden
// rausschmiss aus der Liste nach 20 Sekunden
func register(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("register function is not implemented"))
}
