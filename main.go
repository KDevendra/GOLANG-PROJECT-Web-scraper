package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func main() {

	// Create the CSV file to save the data
	fName := "data.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Could not create file: %v", err)
	}
	defer file.Close()

	// Create a new CSV writer and defer flush
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Initialize a new Colly collector
	c := colly.NewCollector(
		colly.AllowedDomains("udemy.com"),
	)

	// Initialize a counter for the number of occurrences of the specific text
	count := 0

	// Define the specific text you want to count
	desiredText := "desired_text_here"
	// Replace this with the specific text you want to count

	// Define the scraping behavior when an HTML element is found
	c.OnHTML("body", func(e *colly.HTMLElement) {

		// Extract text from the "h1" tag
		text := e.ChildText(".section--section-title--svpHP")
		log.Printf("Scraped text: %s", text)

		// Write the scraped data to the CSV file
		err := writer.Write([]string{text})
		if err != nil {
			log.Printf("Error writing to CSV: %v", err)
		}

		// Check if the scraped text matches the specific text you want to count
		if text == desiredText {
			count++
		}
	})

	// Visit the specified URL
	err = c.Visit("udemy.com/course/go-the-complete-guide")
	if err != nil {
		log.Printf("Error visiting the URL: %v", err)
	}

	// Print completion message and the count of the specific text
	log.Printf("Scraping Completed. Number of occurrences of '%s': %d", desiredText, count)
}
