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
	Image  string `json:"image"`
}

var Movies []Movie

func checkError(error error) {
	if error != nil {
		fmt.Printf("error: %v\n", error)
		log.Fatal(error)
	}
}

func Query(q string) []Movie {

	// Remove whitespace
	q = strings.ReplaceAll(q, " ", "+")

	url := fmt.Sprintf("https://www.imdb.com/find?q=%s&s=tt&ref_=fn_al_tt_mr", q)

	res, err := http.Get(url)
	checkError(err)
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkError(err)

	var imageUrl []string

	doc.Find(".findResult>.primary_photo").Each(func(i int, s *goquery.Selection) {
		if i < 10 {
			itemImg, _ := s.Find("img").Attr("src")
			itemImg = strings.Replace(itemImg, "UX32_CR0,0,32,44", "UX1024_CR0,0,1024,1500", 1) // 1024H x 1500H
			imageUrl = append(imageUrl, itemImg)
		}
	})

	doc.Find(".findResult>.result_text").Each(func(i int, s *goquery.Selection) {
		if i < 10 {
			title := s.Find("a").Text()
			itemUrl, _ := s.Find("a").Attr("href")
			id := strings.Replace(itemUrl, "/title/", "", 1)
			id = strings.Replace(id, "/?ref_=fn_tt_tt_", "", 1)
			id = id[0 : len(id)-1]

			Movies = append(Movies, Movie{Title: title, ImdbId: id, Image: imageUrl[i]})

		}
	})

	return Movies

}
