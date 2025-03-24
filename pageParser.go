package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

// I parse the string as html so i can use the html tags to identify text of interest
// I want to extract the most likely canditates for the jobtitle, for now im assuming it will be the title or the headings.
// Then i extract likely passages of text too breakdown further
func pageParser(htmlString string) (parsedPage, error) {
	possibleTitles := []string{}
	possibleText := []string{}

	htmlReader := strings.NewReader(htmlString)
	doc, err := html.Parse(htmlReader)
	if err != nil {
		return parsedPage{}, fmt.Errorf("failed to parse page: %v", err)
	}
	// Indentify possibleTitles
	for n := range doc.Descendants() {
		if nodeIsPossbileTitle(n) {
			possibleTitles = append(possibleTitles, n.Data)
		}
		if n.Type == html.TextNode {
			if len(n.Data) != 0 {
				possibleText = append(possibleText, n.Data)
			}
		}
	}
	outparsedPage := parsedPage{
		Titles: possibleTitles,
		Text:   possibleText,
	}

	return outparsedPage, nil
}

func nodeIsPossbileTitle(Node *html.Node) bool {
	if Node.Type != html.TextNode {
		return false
	}
	if Node.Parent.Data == "title" || Node.Parent.Data == "h1" || Node.Parent.Data == "h2" || Node.Parent.Data == "h3" || Node.Parent.Data == "h4" || Node.Parent.Data == "h5" || Node.Parent.Data == "h6" {
		return true
	}
	return false
}
