package movieApi

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

var Articles []Article

type Article struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func Get_input() string {
	fmt.Println("enter your query")
	var query string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		query = scanner.Text()
		if strings.Contains(query, " ") {
			query = strings.ReplaceAll(query, " ", "_")
			return query
		} else {
			return query
		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return query
}

func Query_url(q string) {
	fmt.Print(Query(q))
}

func ReturnSingleMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func ReturnQuery(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["q"]

	var queryResult = Query(key) //!TODO : only prints the first query performed
	fmt.Fprintf(w, queryResult[0].Title)
	// fmt.Fprintf(w, url)

}

func ReturnAllMovies(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint: return all articles")
	json.NewEncoder(w).Encode(Articles)
}
