package main

type pageInfo struct {
	Url  string
	Html string
}

type parsedPage struct {
	Titles []string
	Text   []string
}
