package main

import (
	"fmt";
	"regexp";
	"net/http";
	"io";
	"encoding/json";
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

func main() {
	pageURL := "https://apply.workable.com/fuseenergy/j/B73DB96A02/"
	account, shortcode := extractStringRequirements(pageURL)
	apiURL := buildApiURL(account, shortcode)
	bodyHTML := getJSON(apiURL)
	jobData := getParsedData(bodyHTML)
	fmt.Println(account, shortcode)
	fmt.Println(apiURL)
	// fmt.Println(bodyHTML)
	fmt.Printf("%+v\n", jobData)
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
	fmt.Println(m)
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
	}
	defer response.Body.Close()
	// convert JSON to bytes
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
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