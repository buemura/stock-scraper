package cmd

import (
	"fmt"
	"log"
	"strconv"
	"sync"

	"github.com/gocolly/colly"
)

type StockMarketPrice struct {
	Ticker string `json:"ticker"`
	Price float64 `json:"price"`
}

func (c *StockMarketPrice) String() string {
    return fmt.Sprintf(`{"ticker": "%s", "price": %.2f}`, c.Ticker, c.Price)
}

func StartScraping(stocks []string) ([]StockMarketPrice) {
	ch := make(chan StockMarketPrice)
	var wg sync.WaitGroup
	var result []StockMarketPrice

	for _, stock := range stocks {
		wg.Add(1)
		go stockCurrentPrice(stock, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for stockData := range ch {
		result = append(result, stockData)
	}

	return result
}

func stockCurrentPrice(stock string, ch chan StockMarketPrice, wg *sync.WaitGroup) {
	defer (*wg).Done()
	
	c := colly.NewCollector(
	  colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.182 Safari/537.36"),
	  colly.AllowedDomains("finance.yahoo.com"),
	  colly.MaxBodySize(0),
	  colly.AllowURLRevisit(),
	  colly.Async(true),
	)
  
	c.Limit(&colly.LimitRule{
	  DomainGlob:  "*",
	  Parallelism: 2,
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL.String())
	})
  
	c.OnResponse(func(r *colly.Response) {
		log.Println("Received", r.StatusCode)
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Fatal(err.Error())
	})

	c.OnHTML(`fin-streamer[data-field="regularMarketPrice"]`, func(e *colly.HTMLElement) {
	  symbol := e.Attr("data-symbol")
	  if symbol != stock {
		return
	  }
	  
	  value := e.Attr("value")
	  price, _ := strconv.ParseFloat(value, 64)
	  stockMarketPrice := StockMarketPrice{
		Ticker: stock,
		Price: price, 
	  }

	  ch <- stockMarketPrice
	})
  
	c.Visit("https://finance.yahoo.com/quote/"+stock)
	c.Wait()
}