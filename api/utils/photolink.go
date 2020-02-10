package utils

import "github.com/gocolly/colly"

func PhotoLink(isbn string) string {
	c := colly.NewCollector()
	var nulink string
	c.OnHTML(".s-image", func(e *colly.HTMLElement) {
		link := e.Attr("src")
		link = link[36:47]
		nulink = `https://images-na.ssl-images-amazon.com/images/I/` + link + `.jpg`
	})
	c.Visit(`https://www.amazon.com.br/s?k=` + isbn)
	c.Wait()
	return nulink
}
