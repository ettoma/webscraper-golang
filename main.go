package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	// "unicode"

	"github.com/gocolly/colly"
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

func query_url(q string) {
	c := colly.NewCollector(colly.AllowedDomains("en.wikipedia.org"))
	c.OnHTML(".mw-body", func(h *colly.HTMLElement) {
		heading := h.ChildText("h1")
		fmt.Println(heading)
	})
	url := fmt.Sprintf("https://en.wikipedia.org/wiki/%s", q)
	fmt.Println(url)
	c.Visit(url)
}

func main() {

	query_url(get_input())

}
