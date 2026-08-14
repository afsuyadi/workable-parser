package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
)

type jobRecord struct {
	Title       string
	Workplace   string
	JobType     string
	Location    string
	Description string
	PageURL     string
}

// give url and ctx as input, and it will return jobRecord object
func scrapeJob(ctx context.Context, pageURL string) (jobRecord, error) {
	var htmlContent string
	err := chromedp.Run(ctx,
		chromedp.Navigate(pageURL),
		// wait for job title to appear as an indicator
		chromedp.WaitVisible(`[data-ui="job-title"]`, chromedp.ByQuery),
		// give delay time to ensure all elements are written
		chromedp.Sleep(2*time.Second),
		// use OuterHTM to retrieve the outer element of the selector
		chromedp.OuterHTML("html", &htmlContent, chromedp.ByQuery),
	)
	if err != nil {
		return jobRecord{}, err
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		return jobRecord{}, err
	}

	return jobRecord{
		Title:       extractText(doc, `[data-ui="job-title"]`),
		Workplace:   extractText(doc, `[data-ui="job-workplace"]`),
		JobType:     extractText(doc, `[data-ui="job-type"]`),
		Location:    extractText(doc, `[data-ui="job-location-tooltip"]`), // first match = first location, matching the "take the first if there's more than one" rule
		Description: extractText(doc, `[data-ui="job-description"] > div`),
		PageURL:     pageURL,
	}, nil
}

// extract text from the first node matching selector
func extractText(doc *goquery.Document, selector string) string {
	fmt.Println(selector)
	return strings.TrimSpace(doc.Find(selector).First().Text())
}
