package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	// "unicode"

	"github.com/gocolly/colly"
	"github.com/gorilla/mux"
)

func get_input() string {
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

func query_url(q string) string {
	// var heading string
	c := colly.NewCollector(colly.AllowedDomains("en.wikipedia.org"))
	c.OnHTML(".mw-body", func(h *colly.HTMLElement) {
		heading := h.ChildText("h1")
		url := fmt.Sprintf("https://en.wikipedia.org/wiki/%s", q)
		content := h.ChildText("p") // todo: doesn't parse correct element
		newArticle := Article{Id: fmt.Sprint(len(Articles) + 1), Title: heading, Desc: url, Content: content}

		Articles = append(Articles, newArticle)
	},
	)
	url := fmt.Sprintf("https://en.wikipedia.org/wiki/%s", q)
	c.Visit(url)
	return url
}

type Article struct {
	Id      string `json:"id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"Content"`
}

var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the homepage")
	fmt.Println("Endpoint: home page")
}

func handleRequests() {
	port := ":8000"
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", returnAllArticles)
	myRouter.HandleFunc("/articles/{id}", returnSingleArticle)
	myRouter.HandleFunc("/query/{q}", returnQueryResult)
	fmt.Printf("Running on: http://localhost%s \n", port)
	log.Fatal(http.ListenAndServe(port, myRouter))
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint: return all articles")
	json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}

}
func returnQueryResult(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["q"]

	heading := query_url(key)
	fmt.Print(heading)
	fmt.Fprintf(w, heading)

	// fmt.Printf("%s", key)
}

func main() {

	// Articles = []Article{
	// 	{Id: "1", Title: "Hello", Desc: "Ciao", Content: "ok"},
	// 	{Id: "2", Title: "Ok", Desc: "No", Content: "yes"},
	// }
	handleRequests()
	// query_url(get_input())

}
