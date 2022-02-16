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
var imageUrl []string

func checkError(error error) {
	if error != nil {
		fmt.Printf("error: %v\n", error)
		log.Fatal(error)
	}
}

func Query(q string) []Movie {

	//make sure slice is empty
	Movies = nil
	imageUrl = nil

	// Remove whitespace
	q = strings.ReplaceAll(q, " ", "+")

	url := fmt.Sprintf("https://www.imdb.com/find?q=%s&s=tt", q)

	res, err := http.Get(url)
	checkError(err)
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkError(err)

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

func QuerySingleMovie(id string) {

	url := fmt.Sprintf("https://www.imdb.com/title/%s", id)

	res, err := http.Get(url)
	checkError(err)
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkError(err)

	var foundMovie []Movie

	// TODO: doesn't work
	doc.Find("TitleBlock_Container>.TitleBlock__TitleContainer").Each(func(i int, s *goquery.Selection) {

		movieTitle := s.Find("h1").Text()
		fmt.Print(movieTitle)

		foundMovie = append(foundMovie, Movie{Title: movieTitle, ImdbId: id, Image: "ok"})
	})

}
