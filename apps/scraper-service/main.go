package main

import (
	"fmt"
	"scraper-service/cmd"
	"time"
)

func main() {
  start := time.Now()
  stocks := []string {
    "AAPL",
    "MSFT",
    "AMZN",
	"MXRF11.SA",
  }
  
  cmd.StartScraping(stocks)
  fmt.Println("Completed the code process, took: %f seconds", time.Since(start).Seconds())
}
