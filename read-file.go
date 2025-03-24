package main

import (
	"fmt"
	"os"
	"strings"
)

// read entire file based on filepath, split at the first line giving us the url and the html itself.
func ReadFile(filePath string) (pageInfo, error) {
	dat, err := os.ReadFile(filePath)
	if err != nil {
		return pageInfo{}, fmt.Errorf("failed to read file: %v", err)
	}
	url, html, success := strings.Cut(string(dat), "\n")
	if !success {
		return pageInfo{}, fmt.Errorf("failed to split between url and html")
	}
	newPage := pageInfo{
		Url:  url,
		Html: html,
	}
	return newPage, nil
}
