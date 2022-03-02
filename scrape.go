package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"io/ioutil"
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
	c.OnHTML("h5", func(e *colly.HTMLElement) {
        upcomingMissions := &scrapedData {
			Mission: strings.TrimSpace(e.Text),
		}

		// Print mission
        fmt.Printf("Upcoming mission: %s\n", upcomingMissions.Mission)

		// Print mission in JSON format
		missionsJSON, _ := json.Marshal(upcomingMissions)
		fmt.Println(string(missionsJSON))

		// Save mission to JSON file
		file, _ := json.MarshalIndent(upcomingMissions, "", " ")
		_ = ioutil.WriteFile("output.json", file, 0644)
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
	
	
	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})


	// Start scraping on NASA's upcoming launches page
	c.Visit("https://nextspaceflight.com/launches/")
}