package main

import (
	"fmt";
	"regexp";
)



func main() {
	pageURL := "https://apply.workable.com/fuseenergy/j/B73DB96A02/"
	account, shortcode := extractStringRequirements(pageURL)
	fmt.Println(account, shortcode)
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