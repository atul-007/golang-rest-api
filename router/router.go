package router

import (
	controller "github.com/atul-007/golang-rest-api/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controller.Getallmymovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.Createmovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.Markaswatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controller.DeleteAmovie).Methods("DELETE")
	router.HandleFunc("/api/deleteallmovies", controller.DeleteAllMovies).Methods("DELETE")

	return router
}
