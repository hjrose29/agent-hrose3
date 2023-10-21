package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/PuerkitoBio/goquery"
	"strconv"
)

type Datum struct {
	Ticker string
	DateTime time.Time
	Price float32
}

func main() {
	var i int
	stocks := [2]string{"AAPL", "SPY"}

	flag.IntVar(&i, "i", 5, "Time Interval (5 for 5 seconds)")
	flag.Parse()

	ticker := time.NewTicker(time.Duration(i) * time.Second)

		for{
			select {
			case <- ticker.C:

				for _, stock := range stocks{
					
			
					p := getStockPrice(stock)
			
					var datum Datum
					datum.DateTime = time.Now()
					datum.Price = p
					datum.Ticker = stock
				}
			}
			
		}
}

func getStockPrice(Ticker string)(float32) {
	resp, err := http.Get("https://finance.yahoo.com/quote/" + Ticker)
	if(err != nil){
		log.Println(err)
	}
	
	defer resp.Body.Close()
	
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	marketPrice := doc.Find("[data-field=regularMarketPrice][data-symbol=" + Ticker + "]").Text()

	if marketPrice == "" {
		log.Fatalf("Can't Find Price...")
	}

	//I love this format for if statements, it's essentially a try...catch
	if float, err := strconv.ParseFloat(marketPrice, 32); err == nil {
		return float32(float)
	}
	return float32(0)
}