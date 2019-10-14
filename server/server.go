package server

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func logWrapper(f http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Debugf("%s %s", r.Method, r.URL.Path)

		f(w, r)
	})
}

func Start() {
	log.Info("Starting server...")

	r := mux.NewRouter()
	r.HandleFunc("/locate", logWrapper(locateHandler)).Methods("POST")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", http.DefaultServeMux))
}
