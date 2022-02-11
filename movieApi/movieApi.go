package movieApi

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gocolly/colly"
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

func Query_url(q string) string {
	c := colly.NewCollector(colly.AllowedDomains("www.imdb.com"))
	c.OnHTML("#main", func(h *colly.HTMLElement) {
		heading := h.ChildAttr("h1", "class")
		url := fmt.Sprintf("https://www.imdb.com/find?q=%s&ref_=nv_sr_sm", q)
		content := h.ChildText("p") // todo: doesn't parse correct element
		newArticle := Article{Id: fmt.Sprint(len(Articles) + 1), Title: heading, Desc: url, Content: content}

		Articles = append(Articles, newArticle)
	},
	)
	url := fmt.Sprintf("https://www.imdb.com/find?q=%s&ref_=nv_sr_sm", q)
	c.Visit(url)
	return url
}

func ReturnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func ReturnQueryResult(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["q"]

	url := Query_url(key)
	fmt.Print(url)
	fmt.Fprintf(w, url)

}

func ReturnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint: return all articles")
	json.NewEncoder(w).Encode(Articles)
}
