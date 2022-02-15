package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ettoma/web-scraper-go/movieApi"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the homepage")
	fmt.Println("Endpoint: home page")
}

func handleRequests() {
	port := ":8000"
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/movies", movieApi.ReturnAllMovies)
	router.HandleFunc("/movies/{id}", movieApi.ReturnSingleMovie)
	router.HandleFunc("/query/{q}", movieApi.ReturnQuery)
	fmt.Printf("Running on: http://localhost%s \n", port)
	log.Fatal(http.ListenAndServe(port, router))
}

func main() {

	handleRequests()
	// movieApi.Query("the avengers")

}
