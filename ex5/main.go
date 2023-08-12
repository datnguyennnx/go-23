package main

import (
	"sync"

	"github.com/datnguyennnx/go-23/ex5/helper"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	result := make(chan []helper.BodyFetch)
	go helper.Crawl("https://systemdesign.one/categories/", result, &wg)
	go helper.WriteCSV("result.csv", result, &wg)
	wg.Wait()
}
