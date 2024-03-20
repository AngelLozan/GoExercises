package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector()

	url := "https://snapcraft.io/search?q=exodus"

	element := ".p-media-object"

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Printf("Got a response from %v\n\n", r.Request.URL)
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("An error occurred!:", e)
	})

	type Malware struct {
		link string
		title string
	}

	var items []Malware

	c.OnHTML(element, func(e *colly.HTMLElement) {

		maliciousItem := Malware{}

		link := e.Attr("href")
		title := e.Attr("title")

		cleanLink := strings.TrimSpace(link)
		cleanTitle := strings.TrimSpace(title)

		maliciousItem.link = fmt.Sprintf("https://snapcraft.io%s", cleanLink)
		maliciousItem.title = cleanTitle

		if strings.Contains(strings.ToLower(maliciousItem.title), "wallet"){
			items = append(items, maliciousItem)
		}

	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Found some malicious items:", items)
	})

	err := c.Visit(url)
	
	if err != nil {
		log.Fatal(err)
	}
}
