package quotes

import "fmt"

type Pagination struct {
	CurrentPage int `json:"currentPage"`
	Nextpage    int `json:"nextPage"`
	TotalPage   int `json:"totalPages"`
}

type GenresResponse struct {
	Genres []string `json:"data"`
}

type AuthorsResponse struct {
	Authors []string `json:"data"`
}

type QuotaData struct {
	QuoteText   string `json:"quoteText"`
	QuoteAuthor string `json:"quoteAuthor"`
	QuoteGenre  string `json:"quoteGenre"`
}

type QuotesResponse struct {
	Quotas     []QuotaData `json:"data"`
	Pagination Pagination  `json:"pagination"`
}

func QuotesOutput(Quotes []QuotaData) {
	for i, v := range Quotes {
		fmt.Printf("%d. [tag: %s] %s (c) %s\n\n", i, v.QuoteGenre, v.QuoteText, v.QuoteAuthor)
	}
}
