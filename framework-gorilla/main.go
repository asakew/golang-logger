package main

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Infof("Received request: %s %s", r.Method, r.URL)
		next.ServeHTTP(w, r)
		log.Infof("Completed request: %s %s", r.Method, r.URL)
	})
}

func main() {
	router := mux.NewRouter()
	router.Use(loggingMiddleware)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World! ⛅️"))
	})

	log.Info("Starting server on http://localhost:8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}

// time="2025-01-02T11:15:40+10:00" level=info msg="Starting server on http://localhost:8080"
// time="2025-01-02T11:15:48+10:00" level=info msg="Received request: GET /"
