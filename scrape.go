package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

type scrapedData struct {
	Mission string
	// Description string
}

// type jsonData struct {
// 	Mission string
// 	// Description string
// }

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/
func main() {
	// Instantiate default collector
	c := colly.NewCollector(
		colly.CacheDir("./.cache"),
	)
	extensions.RandomUserAgent(c)

	// Mission Title
	c.OnHTML("h3", func(e *colly.HTMLElement) {
		fmt.Print("Hi")
        nasaUpcomingMission := &scrapedData {
			Mission: e.Text,
		}
		

		// Print link
        fmt.Printf("Link found: %q -> %s\n", e.Text, nasaUpcomingMission.Mission)
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})


	
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})
	
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})
	
	c.OnHTML("tr td:nth-of-type(1)", func(e *colly.HTMLElement) {
		fmt.Println("First column of a table row:", e.Text)
	})
	
	c.OnXML("//h1", func(e *colly.XMLElement) {
		fmt.Println(e.Text)
	})
	
	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})


	// Start scraping on NASA's upcoming launches page
	c.Visit("https://www.nasa.gov/launchschedule/")
}