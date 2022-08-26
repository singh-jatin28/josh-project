package routes

import (
	"github.com/gorilla/mux"

	"github.com/singh-jatin28/josh-project/pkg/controllers"
)

var RegisterRoutes = func(router *mux.Router) {
	router.HandleFunc("/websites", controllers.InputSites).Methods("POST")
	router.HandleFunc("/websites", controllers.GetSiteStatus).Methods("GET")
}
