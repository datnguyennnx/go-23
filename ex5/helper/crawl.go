package helper

import (
	"fmt"
	"sync"

	"github.com/gocolly/colly"
)

type BodyFetch struct {
	Url, Title, Subtitle, Date string
}

func Crawl(url string, result chan<- []BodyFetch, wg *sync.WaitGroup) {
	var results = []BodyFetch{}
	defer wg.Done()

	c := colly.NewCollector()
	infoCollector := c.Clone()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	infoCollector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	{
		result := BodyFetch{}
		c.OnHTML("div.card-item > div.card-item-wrapper > article.archive-item ", func(e *colly.HTMLElement) {
			result.Url = e.ChildAttr("a", "href")
			result.Url = e.Request.AbsoluteURL(result.Url)
			infoCollector.Visit(result.Url)
		})
		infoCollector.OnHTML("div.container", func(i *colly.HTMLElement) {
			result.Title = i.ChildText("article.page > h1.single-title")
			result.Subtitle = i.ChildText("article.page > h2.single-subtitle")
			result.Date = i.ChildText("div.post-meta-line > time")
			results = append(results, result)
		})
	}
	c.Visit(url)
	result <- results
	close(result)
}
