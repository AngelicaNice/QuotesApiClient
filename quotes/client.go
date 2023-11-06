package quotes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	client *http.Client
}

func (c Client) Do(req *http.Request) (*http.Response, error) {
	return c.client.Do(req)
}

func NewClient() *Client {
	return &Client{http.DefaultClient}
}

func (c Client) GetGenres() ([]string, error) {
	resp, err := c.client.Get("https://quote-garden.onrender.com/api/v3/Genres")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var r GenresResponse
	if err = json.Unmarshal(body, &r); err != nil {
		return nil, err
	}
	return r.Genres, nil
}

func (c Client) GetAuthors() ([]string, error) {
	resp, err := c.client.Get("https://quote-garden.onrender.com/api/v3/Authors")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var r AuthorsResponse
	if err = json.Unmarshal(body, &r); err != nil {
		return nil, err
	}
	return r.Authors, nil
}

func (c Client) getQuotesByAuthorAndPage(author string, page int) (QuotesResponse, error) {
	req, err := http.NewRequest(http.MethodGet, "https://quote-garden.onrender.com/api/v3/quotes", nil)
	if err != nil {
		return QuotesResponse{}, err
	}

	q := req.URL.Query()
	q.Add("page", fmt.Sprintf("%d", page))
	q.Add("author", author)
	req.URL.RawQuery = q.Encode()

	resp, err := c.Do(req)
	if err != nil {
		return QuotesResponse{}, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var r QuotesResponse
	if err = json.Unmarshal(body, &r); err != nil {
		return QuotesResponse{}, err
	}
	return r, nil
}

func (c Client) GetQuotesByAuthor(author string) ([]QuotaData, error) {
	var quotaData []QuotaData
	page := 1
	for {
		qs, err := c.getQuotesByAuthorAndPage(author, page)
		if err != nil {
			return quotaData, err
		}
		quotaData = append(quotaData, qs.Quotas...)
		if qs.Pagination.TotalPage > page {
			page = page + 1
		} else {
			break
		}
	}
	return quotaData, nil
}
