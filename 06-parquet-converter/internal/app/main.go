package main

import (
	"encoding/json"
	"log"

	"parquet.example/internal/pkg/hotel"
	"parquet.example/internal/pkg/parquet"
)

func main() {
	obj := `[
	{
	  "name": "Grand Hotel Luxor",
	  "city": "Metropolis",
	  "review": 4.3
	},
	{
	  "name": "Sunset Beach Resort",
	  "city": "Seaville",
	  "review": 4.5
	},
	{
	  "name": "Mountain View Lodge",
	  "city": "Peaksville",
	  "review": 4.2
	},
	{
	  "name": "Royal Oasis Palace",
	  "city": "Kingstown",
	  "review": 4.7
	},
	{
	  "name": "Golden Sands Retreat",
	  "city": "Sunnydale",
	  "review": 4.0
	}
  ]
  `
	var hotels []hotel.Hotel
	err := json.Unmarshal([]byte(obj), &hotels)
	if err != nil {
		log.Fatalf("Error parsing json, %v", err)
	}
	err = parquet.StructToParquet("hotels", hotels)
	if err != nil {
		log.Fatalf("Error parsing to parquet, %v", err)
	}
	err = parquet.GenericToParquet("hotelsGeneric", hotels)
	if err != nil {
		log.Fatalf("Error parsing to parquet, %v", err)
	}
}
