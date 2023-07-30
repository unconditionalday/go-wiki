package main

import (
	"fmt"

	gowiki "github.com/unconditionalday/go-wiki/pkg"
)

func main() {
	// Search for the Wikipedia page title
	search_result, _, err := gowiki.Search("verdini", 1, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("This is your search result: %v\n", search_result)

	// Get the page
	page, err := gowiki.GetPage(search_result[0], -1, false, true)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(page.GetSummary())
	fmt.Println(page.GetThumbURL())
	fmt.Println(page.GetLink())
}
