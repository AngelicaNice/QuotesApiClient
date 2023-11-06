package main

import (
	"fmt"

	"github.com/AngelicaNice/QuotesApiClient/quotes"
)

func main() {
	client := quotes.NewClient()
	gs, err := client.GetGenres()
	if err != nil {
		fmt.Println("Genre not found. Please check the query.")
	}
	for i, v := range gs {
		fmt.Printf("%d. %s\n", i, v)
	}

	as, err := client.GetAuthors()
	if err != nil {
		fmt.Println("Something was wrong. Please try again later.")
	}
	for i, v := range as {
		fmt.Printf("%d. %s\n", i, v)
	}

	qs, err := client.GetQuotesByAuthor("Pablo")
	if err != nil {
		fmt.Println("Author not found. Please check the query.")
	}

	quotes.QuotesOutput(qs)
}
