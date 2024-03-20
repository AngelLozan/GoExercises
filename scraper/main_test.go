    package main

		import (
			"testing"
			"github.com/gocolly/colly/v2"
		)
		// URL is not valid or cannot be reached
    // Visits the specified URL and logs the visit
		func TestVisitURLAndLogVisit(t *testing.T) {
			// Initialize the collector
			c := colly.NewCollector()

			// Set up the URL and element to scrape
			url := "https://snapcraft.io/search?q=exodus"
			// element := ".p-media-object"

			// Create a variable to store the visited URL
			var visitedURL string

			// Set up the OnRequest callback to log the visit
			c.OnRequest(func(r *colly.Request) {
				visitedURL = r.URL.String()
			})

			// Visit the URL
			err := c.Visit(url)
			if err != nil {
				t.Fatal(err)
			}

			// Check if the visited URL matches the expected URL
			expectedURL := "https://snapcraft.io/search?q=exodus"
			if visitedURL != expectedURL {
				t.Errorf("Visited URL does not match expected URL. Got: %s, Expected: %s", visitedURL, expectedURL)
			}
		}


		    // URL is invalid or unreachable, logs error and exits
func TestInvalidURLAndLogError(t *testing.T) {
  // Initialize the collector
  c := colly.NewCollector()

  // Set up an invalid URL
  url := "https://invalidurl"

  // Create a variable to store the error message
  var errorMessage string

  // Set up the OnError callback to log the error
  c.OnError(func(r *colly.Response, e error) {
    errorMessage = e.Error()
  })

  // Visit the invalid URL
  err := c.Visit(url)

  // Check if an error occurred
  if err == nil {
    t.Error("Expected an error to occur, but got nil")
  }

  // Check if the error message matches the expected error message
  expectedErrorMessage := "Get \"https://invalidurl\": dial tcp: lookup invalidurl: no such host"
  if errorMessage != expectedErrorMessage {
    t.Errorf("Error message does not match expected error message. Got: %s, Expected: %s", errorMessage, expectedErrorMessage)
  }
}

func TestScrape(t *testing.T) {

  c := colly.NewCollector()

  url := "https://snapcraft.io/search?q=exodus"

  element := ".p-media-object"

	


}
