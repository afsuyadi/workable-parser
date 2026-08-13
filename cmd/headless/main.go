package main

import (
	"fmt";
	"github.com/chromedp/chromedp";
	"context";
	"time";
	"log"
)

func main() {
    // 1. Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// 2. Create the headless browser context
	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	// Variables to store the results
	var pageTitle string
	var bodyText string

	fmt.Println("Launching headless browser and navigating...")

	// 3. Run the automation tasks
	err := chromedp.Run(ctx,
		// a. Navigate to target website
		chromedp.Navigate("https://apply.workable.com/fuseenergy/j/B73DB96A02/"),

		// b. Wait until specific element is visible in the DOM
		chromedp.WaitVisible("body", chromedp.ByQuery),

		// c. Extract the page title
		chromedp.Title(&pageTitle),

		// d. Extract the text inside the <h1> tag
		chromedp.Text("h1", &bodyText, chromedp.ByQuery),
	)
	if err != nil {
		log.Fatalf("Failed to execute browser tasks: %v", err)
	}

	// 4. Output the results to terminal
	fmt.Println("=======Extraction Results============")
	fmt.Printf("Page Title: %s\n", pageTitle)
	fmt.Printf("Heading Text: %s\n", bodyText)
}
