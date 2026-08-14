package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	urlReader "github.com/afsuyadi/workable-parser/pkg/urls-reader"
)

func main() {
	var rows []jobRecord
	urls, err := urlReader.ReadURLs("urls.txt")
	if err != nil {
		fmt.Println(err)
	}
	// get rows for CSV
	for _, url := range urls {
		account, shortcode := extractStringRequirements(url)
		apiURL := buildApiURL(account, shortcode)
		bodyHTML := getJSON(apiURL)
		jobData := getParsedData(bodyHTML)

		workplace := displayWorkplace(jobData.Workplace)
		jobType := getJobType(jobData.JobType)

		row := jobRecord{
			Title:       jobData.Title,
			Workplace:   workplace,
			JobType:     jobType,
			Location:    jobData.Location.City + ", " + jobData.Location.Country,
			Description: jobData.Description,
			PageURL:     url,
		}
		rows = append(rows, row)
	}
	// fmt.Println(rows)

	// get metadata
	fileName := time.Now().Format("2006-01-02") + ".csv"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	writer.Write([]string{"Title", "Workplace", "Job Type", "Location", "Description", "Page URL"})
	for _, row := range rows {
		writer.Write([]string{row.Title, row.Workplace, row.JobType, row.Location, row.Description, row.PageURL})

}	

		
}
