package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"io/ioutil"
	"log"
	"os"
)

type Product struct {
	Url string `json:"url"`
}

func main() {
	products := make([]Product, 0)
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML(".image", func(e *colly.HTMLElement) {
		urlElement := e.DOM.Find("a").Eq(0)

		url, _ := urlElement.Attr("href")

		urls := Product{
			Url: url,
		}
		products = append(products, urls)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	for i := 1; i < 31; i++ {
		a := fmt.Sprintf("https://www.yuzeroglu.com//k//113//okuma-kitaplari?Kategori=113&OrderBy=FiyatArtan&sayfa=%d", i)
		c.Visit(a)
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "	")
		enc.Encode(products)
		writeJSON(products)
	}

}

func writeJSON(data []Product) {
	file, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Println("Unable to create the JSON file.")
	}
	_ = ioutil.WriteFile("urls-data.json", file, 0644)
	fmt.Println("Scraping and Writing successful. Go for Good!")
}
