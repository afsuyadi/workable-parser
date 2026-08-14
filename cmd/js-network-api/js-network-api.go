package main

import (
	"fmt";
	"regexp";
	"net/http";
	"io";
	"encoding/json";
	"encoding/csv";
	"os";
	"time";
	urlReader "github.com/afsuyadi/grab-careers-parser-2/pkg/urls-reader";

)

type WorkableJob struct {
	Title       string `json:"title"`
	Workplace   string `json:"workplace"`
	JobType       string `json:"type"`
	Description string `json:"description"`
	Location    struct {
		City    string `json:"city"`
		Region  string `json:"region"`
		Country string `json:"country"`
	} `json:"location"`
}

type jobRecord struct {
	Title		string
	Workplace	string
	JobType     string
	Location    string
	Description string
	PageURL     string
}

func main() {
	var rows []jobRecord
	urls, err := urlReader.ReadURLs("../../urls.txt")
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

// extract account and shortcode string
func extractStringRequirements(pageURL string) (account string, shortcode string) {
	var workableURLPattern = regexp.MustCompile(`apply\.workable\.com/([^/]+)/j/([^/]+)`)
	m := workableURLPattern.FindStringSubmatch(pageURL)
	if m == nil {
		fmt.Println("URL didn't match the pattern")
	}
	accountId := m[1]
	shortcodeId := m[2]
	// fmt.Println(m)
	return accountId, shortcodeId
}
// ourput: Shortcode and Account

// build API URL using data from network tab
func buildApiURL(account string, shortcode string) (url string) {
	result := fmt.Sprintf("https://apply.workable.com/api/v2/accounts/%s/jobs/%s", account, shortcode)
	return result
}

// get response from API URL 
func getJSON(url string) string {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer response.Body.Close()
	// convert JSON to bytes
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	// convert Bytes to String
	return string(bodyBytes)
}

// get Parsed Data from Body HTML 
func getParsedData(bodyhtml string) WorkableJob {
	var jobData WorkableJob
	// convert raw data to structured value
	err := json.Unmarshal([]byte(bodyhtml), &jobData)
	if err != nil {
		fmt.Println(err)
	}
	return jobData
}

// helper function: display workplace
func displayWorkplace(raw string) string {
	switch raw {
	case "remote": return "Remote"
	case "on_site": return "Onsite"
	case "hybrid": return "Hybrid"
	default: return raw
	}
}


// helper function: job type
func getJobType(jobType string) string {
	switch jobType {
	case "full": return "Full Time"
	default: return jobType
	}
}
