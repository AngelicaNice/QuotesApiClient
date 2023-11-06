# QuotesApiClient
Simple api client for popular quotas getting with Golang from https://pprathameshmore.github.io/QuoteGarden/

## Examlpes:
``` 	
client := quotes.NewClient()
```

For getting all genres:
```
	gs, err := client.GetGenres()
	if err != nil {
		fmt.Println("Genre not found. Please check the query.")
	}
	for i, v := range gs {
		fmt.Printf("%d. %s\n", i, v)
	}
```

For getting all authors:
```
	as, err := client.GetAuthors()
	if err != nil {
		fmt.Println("Something was wrong. Please try again later.")
	}
	for i, v := range as {
		fmt.Printf("%d. %s\n", i, v)
	}
```

For getting all quotes by author:
```
	qs, err := client.GetQuotesByAuthor("Pablo")

	if err != nil {
		fmt.Println("Author not found. Please check the query.")
	}

	quotes.QuotesOutput(qs)
```