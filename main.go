package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	// Create a new collector
	c := colly.NewCollector()

	// Callback function to be executed when a link is found
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		// Extract the URL from the anchor element and visit it
		e.Request.Visit(e.Attr("href"))
	})

	// Callback function to be executed before making an HTTP request
	c.OnRequest(func(r *colly.Request) {
		// Print the URL being visited
		fmt.Println("Visiting", r.URL)
	})

	// Start the crawling process by visiting the initial URL
	initialURL := "http://go-colly.org/"
	c.Visit(initialURL)
}
