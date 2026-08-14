package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/chromedp/cdproto/cdp"
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
		chromedp.OuterHTML("html", &htmlContent, chromedp.ByQuery),
		// disini langsung chrome.Run, gaperlu di ExtractText()
	)
	if err != nil {
		return jobRecord{}, err
	}
	// fmt.Println(extractText(ctx, `[data-ui="job-title"]`))
	return jobRecord{
		Title:       extractText(ctx, `[data-ui="job-title"]`),
		Workplace:   extractText(ctx, `[data-ui="job-workplace"]`),
		JobType:     extractText(ctx, `[data-ui="job-type"]`),
		Location:    extractText(ctx, `[data-ui="job-location-tooltip"]`), // first match = first location, matching the "take the first if there's more than one" rule
		Description: extractText(ctx, `[data-ui="job-description"] > div`),
		PageURL:     pageURL,
	}, nil
}

// extract text from a html node
func extractText(ctx context.Context, selector string) string {
	var nodes []*cdp.Node
	if err := chromedp.Run(ctx, chromedp.Nodes(selector, &nodes, chromedp.ByQuery, chromedp.AtLeast(0))); err != nil {
		return ""
	}
	if len(nodes) == 0 {
		return ""
	}

	var text string
	if err := chromedp.Run(ctx, chromedp.Text(selector, &text, chromedp.ByQuery)); err != nil {
		return ""
	}
	fmt.Println(text)
	return strings.TrimSpace(text)
}
