package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"strconv"
	"strings"
)

func main()  {
	fmt.Println(GetPrice())

}

func GetPrice() float64 {
	c := colly.NewCollector()

	// Find and visit all links
	var price float64
	c.OnHTML(".priceValue___11gHJ", func(e *colly.HTMLElement) {
		ps := strings.ReplaceAll(e.Text, "$0.", "")
		float, err := strconv.ParseFloat(ps, 64)
		if err != nil {
			fmt.Println(err)
		}
		price = float
	})

	c.Visit("https://coinmarketcap.com/currencies/safemoon/")

	return price
}