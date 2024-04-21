package main

import (
	"fmt"
	"log"

	"parquet.example/internal/pkg/parquet"
)

func main() {
	// 	obj := `[
	// 	{
	// 	  "name": "Grand Hotel Luxor",
	// 	  "city": "Metropolis",
	// 	  "review": 4.3
	// 	},
	// 	{
	// 	  "name": "Sunset Beach Resort",
	// 	  "city": "Seaville",
	// 	  "review": 4.5
	// 	},
	// 	{
	// 	  "name": "Mountain View Lodge",
	// 	  "city": "Peaksville",
	// 	  "review": 4.2
	// 	},
	// 	{
	// 	  "name": "Royal Oasis Palace",
	// 	  "city": "Kingstown",
	// 	  "review": 4.7
	// 	},
	// 	{
	// 	  "name": "Golden Sands Retreat",
	// 	  "city": "Sunnydale",
	// 	  "review": 4.0
	// 	}
	//   ]
	//   `
	// var hotels []hotel.Hotel
	// err := json.Unmarshal([]byte(obj), &hotels)
	// if err != nil {
	// 	log.Fatalf("Error parsing json, %v", err)
	// }
	// var names []string
	// var cities []string
	// var reviews []float32
	// for _, item := range hotels {
	// 	names = append(names, item.Name)
	// 	cities = append(cities, item.City)
	// 	reviews = append(reviews, item.Review)
	// }

	// parquet.BuildStruct(names, cities, reviews)
	// parquet.ReadCSV()
	// currentDir, err := os.Getwd()
	// if err != nil {
	// 	fmt.Println("Erro ao obter o diretório de trabalho atual:", err)
	// 	return
	// }
	// fmt.Println("Diretório de trabalho atual:", currentDir)

	// parquet.JsonFileToParquet("hotels_100000.json")
	// parquet.CsvToParquet("hotels_10000.csv")
	chan1 := parquet.CsvToArrowChannel("hotels_10000.csv")
	fmt.Println(chan1)
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
	chan8, errChan := parquet.JsonFileToArrowChannel("hotels_10000_copy.json")
	if errChan != nil {
		for err := range errChan {
			log.Println(err)
		}
	}
	fmt.Println(chan8)
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
	for c := range parquet.FanIn(chan1, chan8) {
		parquet.ArrowToParquet(c)
		// fmt.Println(c)
	}
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
