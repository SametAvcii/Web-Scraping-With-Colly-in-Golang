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

type Product_Book struct {
	Name   string `json:"name"`
	Barkod string `json:"barkod"`
	Price  string `json:"price"`
	Brand  string `json:"brand:"`
}

func main() {

	c := colly.NewCollector()

	product_books := make([]Product_Book, 0)
	c.OnHTML("div.product>div.row", func(k *colly.HTMLElement) {

		Barkod := k.ChildText("div.col-xs-12.col-md-8.col-lg-8.product-detail-col>div.product-detail>ul.info.hidden-xs.hidden-sm>li > span#kod")
		Name := k.ChildText("div.col-xs-12.col-md-8.col-lg-8.product-detail-col>div.product-detail>h2#baslik")
		Brand := k.ChildText("div.col-xs-12.col-md-8.col-lg-8.product-detail-col>div.product-detail>a>div.text")

		Check := k.ChildText("div.col-xs-12.col-md-4.col-lg-4.product-cart-col>div.product-cart>div.price>div#satis-fiyati>span#satis")
		if Check == "" {
			Check = k.ChildText("div.col-xs-12.col-md-4.col-lg-4.product-cart-col>div.product-cart>div.price>div#indirimli-fiyat>span#indirimli")
		}

		products := Product_Book{
			Name:   Name,
			Barkod: Barkod,
			Brand:  Brand,
			Price:  Check,
		}
		product_books = append(product_books, products)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	jsonUrl, err := os.Open("urls-data.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonUrl.Close()

	urls, err := ioutil.ReadAll(jsonUrl)
	if err != nil {
		fmt.Println(err)
	}
	var product []Product
	json.Unmarshal(urls, &product)
	for _, urlArr := range product {

		c.Visit(urlArr.Url)

	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "	")
	enc.Encode(product_books)
	writeJSON2(product_books)
}
func writeJSON2(data []Product_Book) {
	file, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Println("Unable to create the JSON file.")
	}
	_ = ioutil.WriteFile("barkods-data.json", file, 0644)
	fmt.Println("Barkods Append. Go for Good!")

}
