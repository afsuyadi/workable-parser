package io

import (
	"bufio"
	"os"
)

func ReadURLs(path string) ([]string, error){
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var urls []string
	url := bufio.NewScanner(f)
	for url.Scan() {
		urlText := url.Text()
		if urlText == "" {
			continue
		}
		urls = append(urls, urlText)
	}
	return urls, nil
}