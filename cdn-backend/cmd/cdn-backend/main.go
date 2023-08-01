package main

import (
	"log"
	"net/http"
	"github.com/ra/cdn-backend/api/handler"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/upload", handler.UploadFile).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}
