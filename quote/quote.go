package quote

import (
	"encoding/json"
	"fmt"
	"net/http"

	util "github.com/mikeunge/Happy-Fetch/utils"
)

type Quote struct {
	Text   string
	Author string
}

// GetQuoteFromQuotes :: get a single (random) quote.
func GetQuoteFromQuotes(quotes []Quote) Quote {
	var q Quote
	choice := util.RandInt(0, len(quotes))
	q = quotes[choice]
	return q
}

// GetQuotesFromApi :: fetch the API for it's quotes.
func GetQuotesFromApi(apiEndpoint string) ([]Quote, error) {
	var quotes []Quote

	r, err := http.Get(apiEndpoint)
	if err != nil {
		return quotes, err
	}
	err = json.NewDecoder(r.Body).Decode(&quotes)
	if err != nil {
		return quotes, err
	}

	return quotes, nil
}

// FormatQuote :: return a (formatted) string containing the quote Text and the Author.
func FormatQuote(q Quote) (string, error) {
	if len(q.Text) == 0 {
		return "", fmt.Errorf("no text to output")
	} else if len(q.Author) == 0 {
		fmt.Println("WARN: no author to display")
		return q.Text, nil
	} else {
		return fmt.Sprintf("%s -%s", q.Text, q.Author), nil
	}
}
