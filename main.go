package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/pecan-pie/zcpm-central/handler"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

func main() {
	var log = logrus.New()

	// set the log output
	log.Out = os.Stdout

	// set the log level
	log.Level = logrus.DebugLevel

	router := mux.NewRouter()

	router.Path("/metrics").Handler(promhttp.Handler())
	router.Path("/register").Handler(handler.CreateInstrumentedHandler(handler.NewRegisterHandler(), "register"))
	router.Path("/random").Handler(handler.CreateInstrumentedHandler(handler.NewRandomHandler(), "random"))
	router.Path("/validate").Handler(handler.CreateInstrumentedHandler(handler.NewValidateHandler(), "validate"))

	log.Fatal(http.ListenAndServe(":8000", router))
}
