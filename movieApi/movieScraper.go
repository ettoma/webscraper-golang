package movieApi

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Movie struct {
	ImdbId string `json:"imdbId"`
	Title  string `json:"title"`
}

var Movies []Movie

func checkError(error error) {
	if error != nil {
		fmt.Printf("error: %v\n", error)
		log.Fatal(error)
	}
}

func Query(q string) {

	// Remove whitespace
	q = strings.ReplaceAll(q, " ", "+")

	url := fmt.Sprintf("https://www.imdb.com/find?q=%s&s=tt&ref_=fn_al_tt_mr", q)

	res, err := http.Get(url)
	checkError(err)
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkError(err)

	doc.Find(".findResult>.result_text").Each(func(i int, s *goquery.Selection) {
		if i < 10 {
			title := s.Find("a").Text()
			// img, _ := s.Find("img").Attr("src") //Todo: implement IMG src parser
			itemUrl, _ := s.Find("a").Attr("href")
			id := strings.Replace(itemUrl, "/title/", "", 1)
			id = strings.Replace(id, "/?ref_=fn_tt_tt_", "", 1)
			id = id[0 : len(id)-1]

			// fmt.Print(img)

			Movies = append(Movies, Movie{Title: title, ImdbId: id})

		}
	})
}
