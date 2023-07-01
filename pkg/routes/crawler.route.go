package routes

import (
	"github.com/2O23/crawler/pkg/controllers"
	"github.com/gorilla/mux"
)

var CrawlerRoutes = func(router *mux.Router) {
	router.HandleFunc("/sites",controllers.GetSites).Methods("GET")
	router.HandleFunc("/sites",controllers.CreateSite).Methods("POST")//TODO:create controllers
	router.HandleFunc("/files",controllers.GetFiles).Methods("GET")
	router.HandleFunc("/sites/{id}",controllers.UpdateFile).Methods("PUT")
	router.HandleFunc("/files/{id}",controllers.UpdateFile).Methods("PUT")
}