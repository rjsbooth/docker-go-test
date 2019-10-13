package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.Path("/api/v0/health").Methods("GET").HandlerFunc(healthcheckHandler)
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Test service is ok"))
	w.WriteHeader(http.StatusOK)
}
