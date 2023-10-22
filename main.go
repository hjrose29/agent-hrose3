package main

import (
	"flag"
	"log"
	"net/http"
	"time"
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/jamespearly/loggly"
)

type Datum struct {
	Ticker string
	DateTime int64
	Price float32
}

func main() {
	var i int

	client := loggly.New("Web Scraper")
	
	stocks := [4]string{"AAPL", "SPY", "AMZN", "TSLA"}

	flag.IntVar(&i, "i", 60, "Time Interval (5 for 5 seconds)")
	flag.Parse()

	ticker := time.NewTicker(time.Duration(i) * time.Second)

		for{
			select {
			case <- ticker.C:

				for _, stock := range stocks{
					
			
					p := getStockPrice(stock)			
					
					var datum Datum
					datum.DateTime = time.Now().Unix()
					datum.Price = p
					datum.Ticker = stock
					dynamoDBWrite(datum)
					err := client.EchoSend("info", "Wrote Data for " + datum.Ticker)
 					if(err != nil){fmt.Println("err:", err)}
	
				}
			}
			
		}
}

func getStockPrice(Ticker string)(float32) {
	client := loggly.New("Web Scraper")
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
		client.Send("error", "Can't find price")
	}

	//I love this format for if statements, it's essentially a try...catch
	if float, err := strconv.ParseFloat(marketPrice, 32); err == nil {
		return float32(float)
	}
	return float32(0)
}
func dynamoDBWrite(datum Datum){
	client := loggly.New("Web Scraper")
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := dynamodb.New(sess)
	av, err := dynamodbattribute.MarshalMap(datum)
	if err != nil {
		log.Fatalf("Error formatting data: %v", err)
		client.Send("error", "Error formatting data")
	}
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("hrose3"),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		log.Fatalf("Error calling PutItem: %v", err)
		client.Send("error", "Error calling PutItem")
	}

}