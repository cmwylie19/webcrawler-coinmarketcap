package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

type currency struct {
	Name      string
	Symbol    string
	Price     string
	Volume    string
	MarketCap string
}

func main() {
	// Instantiate default collector
	c := colly.NewCollector()
	messages := []currency{}

	c.OnHTML("tbody tr", func(e *colly.HTMLElement) {

		messages = append(messages, currency{
			Name:      strings.Split(e.ChildText(".cmc-link"), "$")[0],
			Symbol:    e.ChildText("td.cmc-table__cell--sort-by__symbol"),
			Price:     e.ChildText("td.cmc-table__cell--sort-by__price"),
			Volume:    e.ChildText("td.cmc-table__cell--sort-by__volume-24-h"),
			MarketCap: e.ChildText(".sc-1eb5slv-0"),
		})
	})

	err := c.Visit("https://coinmarketcap.com/all/views/all/")
	if err != nil {
		panic(err)
	}

	c.Wait()

	bs, err := json.MarshalIndent(messages, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bs))
	fmt.Println("Crypto currencies scraped:", len(messages))
}
