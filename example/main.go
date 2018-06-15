package main

import (
	"github.com/kr/pretty"
	"github.com/susloparovdenis/scrapping"
)

func main() {
	engines := []scrapping.SearchEngine{scrapping.NewGoogle(), scrapping.NewYahoo()}
	aggregator := scrapping.Aggregator{engines}
	results := aggregator.Query("concurrency wiki")
	pretty.Print(results)
}
