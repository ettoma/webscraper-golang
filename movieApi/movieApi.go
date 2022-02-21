package movieApi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the homepage")
	fmt.Println("Endpoint: home page")
}

func ReturnSingleMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	w.Header().Set("Content-Type", "application/json")

	var movieData = QuerySingleMovie(id)

	json.NewEncoder(w).Encode(movieData)

}

func ReturnMoviesFromQuery(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["q"]

	w.Header().Set("Content-Type", "application/json")

	var queryResults = QueryAllMovies(key)

	json.NewEncoder(w).Encode(queryResults)

}

func ReturnAllMovies(w http.ResponseWriter, r *http.Request) {

}

func PostSingleMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	fmt.Print(QuerySingleMovie(vars["id"]))
}
