package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"

	urlReader "github.com/afsuyadi/grab-careers-parser-2/pkg/urls-reader"
)

type jobRecord struct {
	Title       string
	Workplace   string
	JobType     string
	Location    string
	Description string
	PageURL     string
}

func main() {
	urls, err := urlReader.ReadURLs("urls.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// start a browser instance
	browserCtx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// create a ctx that will be reused
	ctx, cancel := context.WithTimeout(browserCtx, 3*time.Minute)
	defer cancel()

	var rows []jobRecord
	// read every url and save it into the array
	for _, url := range urls {
		row, err := scrapeJob(ctx, url)
		if err != nil {
			fmt.Println("skipping", url, "-", err)
			continue
		}
		rows = append(rows, row)
	}

	// create the csv file with rows array data
	fileName := time.Now().Format("2006-01-02") + ".csv"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// create the columns of csv
	writer := csv.NewWriter(file)
	defer writer.Flush()
	writer.Write([]string{"Title", "Workplace", "Job Type", "Location", "Description", "Page URL"})
	for _, row := range rows {
		writer.Write([]string{row.Title, row.Workplace, row.JobType, row.Location, row.Description, row.PageURL})
	}

	fmt.Printf("wrote %d rows to %s\n", len(rows), fileName)
}

// give url and ctx as input, and it will return jobRecord object
func scrapeJob(ctx context.Context, pageURL string) (jobRecord, error) {
	err := chromedp.Run(ctx,
		chromedp.Navigate(pageURL),
		// wait for job title to appear as an indicator
		chromedp.WaitVisible(`[data-ui="job-title"]`, chromedp.ByQuery),
		chromedp.Sleep(2*time.Second),
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
