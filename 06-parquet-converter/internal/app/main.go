package main

import (
	"fmt"
	"os"
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
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Erro ao obter o diretório de trabalho atual:", err)
		return
	}
	fmt.Println("Diretório de trabalho atual:", currentDir)
	// parquet.JsonFileToParquet("hotels.json", utils.GetParquetFilePath())
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
