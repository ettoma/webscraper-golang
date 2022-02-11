package movieApi

import (
	"fmt"

	"github.com/gocolly/colly"
)

func Query(q string) {
	c := colly.NewCollector(colly.AllowedDomains("www.imdb.com"))

	c.OnHTML("#main", func(h *colly.HTMLElement) {
		text := h.ChildAttr("h1", "class")
		fmt.Print(text)

	})
	url := fmt.Sprintf("https://www.imdb.com/find?q=%s&ref_=nv_sr_sm", q)
	c.Visit(url)

	// c := colly.NewCollector(colly.AllowedDomains("www.imdb.com"))
	// c.OnHTML("#main", func(h *colly.HTMLElement) {
	// 	heading := h.ChildAttr("h1", "class")
	// 	url := fmt.Sprintf("https://www.imdb.com/find?q=%s&ref_=nv_sr_sm", q)
	// 	content := h.ChildText("p") // todo: doesn't parse correct element
	// 	newArticle := Article{Id: fmt.Sprint(len(Articles) + 1), Title: heading, Desc: url, Content: content}

	// 	Articles = append(Articles, newArticle)
	// },
	// )
	// url := fmt.Sprintf("https://www.imdb.com/find?q=%s&ref_=nv_sr_sm", q)
	// c.Visit(url)
	// return url
}
