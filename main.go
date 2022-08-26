package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/singh-jatin28/josh-project/pkg/routes"
)

func main() {

	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	http.Handle("/", r)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
