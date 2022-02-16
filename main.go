package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ettoma/web-scraper-go/movieApi"

	"github.com/gorilla/mux"
)

func handleRequests() {
	port := ":8000"
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", movieApi.HomePage)
	router.HandleFunc("/movies", movieApi.ReturnAllMovies)
	router.HandleFunc("/movies/id={id}", movieApi.ReturnSingleMovie)
	router.HandleFunc("/movies/q={q}", movieApi.ReturnMoviesFromQuery)
	fmt.Printf("Running on: http://localhost%s \n", port)
	log.Fatal(http.ListenAndServe(port, router))
}

func main() {

	handleRequests()
	// movieApi.Query("the avengers")

}
