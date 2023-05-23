package controller

import (
	"encoding/json"
	"net/http"

	"github.com/atul-007/golang-rest-api/helper"
	"github.com/atul-007/golang-rest-api/models"
	"github.com/gorilla/mux"
)

func Getallmymovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allmovies := helper.Getallmovies()
	json.NewEncoder(w).Encode(allmovies)
}

func Createmovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Method", "POST")

	var movie models.Netflix
	json.NewDecoder(r.Body).Decode(&movie)
	helper.Insertonemovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func Markaswatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Method", "POST")

	params := mux.Vars(r)
	helper.Updateonerecord(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAmovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Method", "DELETE")

	params := mux.Vars(r)
	helper.Deleteonerecord(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Method", "DELETE")

	count := helper.Deleteallrecords()
	json.NewEncoder(w).Encode(count)
}
