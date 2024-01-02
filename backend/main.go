package main

import (
	"bookmark-backend/router"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func MiddlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token")

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(MiddlewareCors)

	router.NewFolderRouter("/folder", r)
	router.NewFileRouter("/file", r)

	log.Println("Connected to port: 8000")

	http.ListenAndServe(":8000", r)
}
