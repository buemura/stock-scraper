package main

import (
	"fmt"
	"scraper-service/cmd"
	"time"
)

func main() {
  start := time.Now()
  stocks := []string {
    "VTI",
    "PFF",
    "JEPI",
    "PGX",
	"LQD",
	"BCFF11.SA",
	"BDIF11.SA",
	"BRCR11.SA",
	"BTLG11.SA",
	"HCTR11.SA",
	"HGRE11.SA",
	"IFRA11.SA",
	"KISU11.SA",
	"MXRF11.SA",
	"RZAG11.SA",
	"VISC11.SA",
	"XPCA11.SA",
	"XPLG11.SA",
	"XPML11.SA",
  }
  
  cmd.StartScraping(stocks)
  fmt.Println("Completed the code process, took: %f seconds", time.Since(start).Seconds())
}
