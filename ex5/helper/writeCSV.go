package helper

import (
	"encoding/csv"
	"errors"
	"os"
	"sync"
)

func WriteCSV(nameCSV string, results <-chan []BodyFetch, wg *sync.WaitGroup) {
	defer wg.Done()

	file, err := os.Create(nameCSV)
	if err != nil {
		panic(errors.New("failed to create output CSV file"))

	}
	defer file.Close()
	writer := csv.NewWriter(file)
	writer.Write([]string{"URL", "Title", "Subtitle", "Date"})
	for rs := range results {
		for _, value := range rs {
			writer.Write([]string{value.Url, value.Title, value.Subtitle, value.Date})
		}
		_, ok := <-results
		if !ok {
			break
		}

	}
	defer writer.Flush()
}
