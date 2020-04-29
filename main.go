package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"testpr/pkg/api"
)
func main() {
	r := mux.NewRouter()
	api.RegisterHandlers(r)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}