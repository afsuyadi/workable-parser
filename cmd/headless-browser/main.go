package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"github.com/chromedp/chromedp"

	urlReader "github.com/afsuyadi/workable-parser/pkg/urls-reader"
)

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

