package main

import (
	"fmt"
	"sync"

	"github.com/apache/arrow/go/v16/arrow"
	"parquet.example/internal/pkg/parquet"
)

func main() {
	//fazer uso de um channel e executar as funçoes com Go
	//utilizar wait para esperar o termino das goroutines e retornar com mensagem
	chanIn := make(chan []arrow.Record)
	var wg = &sync.WaitGroup{}

	// concurrency, err := strconv.Atoi(os.Getenv("CONCURRENCY_WORKERS"))
	// if err != nil {
	// 	log.Fatalf("error loading var: CONCURRENCY_WORKERS.")
	// }
	concurrency := 1

	for qtdProcessesCsv := 0; qtdProcessesCsv < concurrency; qtdProcessesCsv++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			parquet.CsvToArrowChannel("hotels_10000.csv", chanIn)
		}()

	}
	fmt.Println(chanIn)
	for qtdProcessesJson := 0; qtdProcessesJson < concurrency; qtdProcessesJson++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			parquet.JsonFileToArrowChannel("hotels_10000_copy.json", chanIn)
		}()

	}

	go func() {
		wg.Wait()
		close(chanIn)
	}()

	for c := range chanIn {
		parquet.ArrowToParquet(c)
		fmt.Println(c)
	}
	// chan2 := parquet.CsvToArrowChannel("hotels_10000_copy.csv")
	// chan3 := parquet.CsvToArrowChannel("hotels_10000_copy_2.csv")
	// chan4 := parquet.CsvToArrowChannel("hotels_10000_copy_3.csv")
	// chan5 := parquet.CsvToArrowChannel("hotels_10000_copy_4.csv")
	// chan6 := parquet.CsvToArrowChannel("hotels_10000_copy_5.csv")
	// chan7, errChan := parquet.JsonFileToArrowChannel("hotels_10000.json")
	// if errChan != nil {
	// 	for err := range errChan {
	// 		log.Println(err)
	// 	}
	// }

	//fazer uso de um channel e executar as funçoes com Go
	//utilizar wait para esperar o termino das goroutines e retornar com mensagem

	// chan9, errChan := parquet.JsonFileToArrowChannel("hotels_10000_copy_2.json")
	// if errChan != nil {
	// 	for err := range errChan {
	// 		log.Println(err)
	// 	}
	// }
	// chan10, errChan := parquet.JsonFileToArrowChannel("hotels_10000_copy_3.json")
	// if errChan != nil {
	// 	for err := range errChan {
	// 		log.Println(err)
	// 	}
	// }
	// chan11, errChan := parquet.JsonFileToArrowChannel("hotels_10000_copy_4.json")
	// if errChan != nil {
	// 	for err := range errChan {
	// 		log.Println(err)
	// 	}
	// }
	// chan12, errChan := parquet.JsonFileToArrowChannel("hotels_10000_copy_5.json")
	// if errChan != nil {
	// 	for err := range errChan {
	// 		log.Println(err)
	// 	}
	// }

	//fazer uso de um channel e executar as funçoes com Go

	// a :=[]chan
	// for c := range parquet.FanIn(chan1, chan2, chan3, chan4, chan5, chan6, chan7, chan8, chan9, chan10, chan11, chan12) {
	// 	parquet.ArrowToParquet(c)
	// 	fmt.Println(c)
	// }
	// numRecords := 10000

	// // Gerar registros
	// records := utils.GenerateRecords(numRecords)

	// // Escrever registros em um arquivo CSV
	// filename := "records.csv"
	// if err := utils.WriteCSV(filename, records); err != nil {
	// 	log.Fatalf("Erro ao escrever CSV: %v", err)
	// }
	// log.Printf("Arquivo CSV com %d registros criado com sucesso: %s", numRecords, filename)

	// utils.GenerateJson()
	// parquet.CsvFileToParquet()
	// parquet.JsonToParquet()
	// parquet.CsvToParquet()

	// err = parquet.StructToParquet("hotels", hotels)
	// if err != nil {
	// 	log.Fatalf("Error parsing to parquet, %v", err)
	// }
	// err = parquet.GenericToParquet("hotelsGeneric", hotels)
	// if err != nil {
	// 	log.Fatalf("Error parsing to parquet, %v", err)
	// }

}
