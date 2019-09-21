package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {

	mux := mux.NewRouter()

	mux.HandleFunc("/health", middleware(health)).Methods("GET")
	mux.HandleFunc("/track", middleware(tracking)).Methods("POST")
	mux.HandleFunc("/search", middleware(search)).Methods("POST")

	return mux
}

// General middleware with basic logging
func middleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
		log.Println(r.Method, r.RequestURI, time.Since(startTime))
	})
}
