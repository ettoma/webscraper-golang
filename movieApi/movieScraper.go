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

type MovieDetails struct {
	ImdbId     string `json:"imdbId"`
	Title      string `json:"title"`
	Year       string `json:"year"`
	Duration   string `json:"duration"`
	ImdbRating string `json:"rating"`
}

func checkError(error error) {
	if error != nil {
		log.Fatalln("Error: ", error)
	}
}

func QueryAllMovies(q string) []Movie {
	var Movies []Movie

	// Remove whitespace from query if any
	q = strings.ReplaceAll(q, " ", "+")

	url := fmt.Sprintf("https://www.imdb.com/find?q=%s&s=tt", q)

	res, err := http.Get(url)
	checkError(err)
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkError(err)

	doc.Find(".findResult").Each(func(i int, s *goquery.Selection) {
		if i < 10 {
			title := s.Find(".result_text>a").Text()

			itemUrl, _ := s.Find(".result_text>a").Attr("href")

			id := strings.Replace(itemUrl, "/title/", "", 1)
			id = strings.Replace(id, "/?ref_=fn_tt_tt_", "", 1)
			id = id[0 : len(id)-1]

			itemImgContainer := s.Find(".primary_photo")
			itemImg, _ := itemImgContainer.Find("img").Attr("src")
			itemImg = strings.Replace(itemImg, "UX32_CR0,0,32,44", "UX1024_CR0,0,1024,1500", 1)

			Movies = append(Movies, Movie{Title: title, ImdbId: id, Image: itemImg})

		}
	})

	return Movies

}

func QuerySingleMovie(id string) MovieDetails {

	var movieTitle string
	var yearText string
	var durationText string
	var foundRating string

	url := fmt.Sprintf("https://www.imdb.com/title/%s", id)

	res, err := http.Get(url)
	checkError(err)
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkError(err)

	doc.Find(".RatingBar__RatingContainer-sc-85l9wd-0").Each(func(i int, s *goquery.Selection) {

		foundRating = s.Find("span").First().Text()
		if len(foundRating) != 0 {

			foundRating = foundRating[0:3]
		}

	})

	doc.Find(".TitleBlock__Container-sc-1nlhx7j-0>.TitleBlock__TitleContainer-sc-1nlhx7j-1").Each(func(i int, s *goquery.Selection) {

		movieTitle = s.Find("h1").Text()
		year := s.Find("span")
		yearText = year.First().Text()

		duration := s.Find("li")
		durationText = duration.Last().Text()

	})

	doc.Find(".Storyline__StorylineWrapper-sc-1b58ttw-0").Each(func(i int, s *goquery.Selection) {
		// fmt.Print(s.Find("div").First().Text())
	})

	return MovieDetails{Title: movieTitle, ImdbId: id, Year: yearText, Duration: durationText, ImdbRating: foundRating}

}
