package main

import (
	"log"
	"sync"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/joho/godotenv"
	"parquet.example/internal/pkg/parquet"
)

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

}

func main() {
	chanIn := make(chan []arrow.Record, 20)
	var wg = &sync.WaitGroup{}
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	parquet.CsvToArrowChannel("hotels_data.csv", chanIn)
	// }()

	wg.Add(1)
	go func() {
		defer wg.Done()
		parquet.JsonFileToArrowChannel("hotels_5.json", chanIn)
	}()

	wg.Wait()
	close(chanIn)

	for record := range chanIn {
		parquet.ArrowToParquet(record)
	}

}
