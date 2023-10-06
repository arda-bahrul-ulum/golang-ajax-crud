package main

import (
	"golang-ajax-crud/controllers/mahasiswacontroller"
	"net/http"
)

func main() {
	http.HandleFunc("/", mahasiswacontroller.Index)
	http.HandleFunc("/mahasiswa/get_form", mahasiswacontroller.GetForm)
	http.HandleFunc("/mahasiswa/store", mahasiswacontroller.Store)
	http.HandleFunc("/mahasiswa/delete", mahasiswacontroller.Delete)

	http.ListenAndServe(":4000", nil)
}
