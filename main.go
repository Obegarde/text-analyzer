package main

import "fmt"

func main() {
	testFilePath := "/home/obegarde/workspace/github.com/obegarde/crawler/out/www.jobindex.dk/008bfc327331c7ec1da95fa22a5588f7e179fa44661e01d42fbb4ca852d40bc4"

	pageInfo, err := ReadFile(testFilePath)
	if err != nil {
		fmt.Printf("Error reading pageInfo %v\n", err)
	}
	fmt.Println(pageInfo.Url)
	pPage, err := pageParser(pageInfo.Html)
	if err != nil {
		fmt.Println(err)
	}
	for _, val := range pPage.Titles {
		fmt.Println(val)
	}
	for _, val := range pPage.Text {
		fmt.Printf("%v", val)
	}
}
