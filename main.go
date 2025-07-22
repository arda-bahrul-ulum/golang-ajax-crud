package main

import (
	"golang-ajax-crud/controllers/mahasiswacontroller"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", mahasiswacontroller.Index)
	mux.HandleFunc("/mahasiswa/get_form", mahasiswacontroller.GetForm)
	mux.HandleFunc("/mahasiswa/store", mahasiswacontroller.Store)
	mux.HandleFunc("/mahasiswa/delete", mahasiswacontroller.Delete)

	// CORS Middleware
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow any origin
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Serve the actual request
		mux.ServeHTTP(w, r)
	})

	http.ListenAndServe(":4000", handler)
}