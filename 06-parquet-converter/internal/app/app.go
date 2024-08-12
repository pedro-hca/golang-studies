package main

import (
	"encoding/json"

	"parquet.example/internal/pkg/hotel"
	"parquet.example/internal/pkg/parquet"
	"parquet.example/internal/pkg/schema"
)

// func init() {
// 	err := godotenv.Load("../../.env")
// 	if err != nil {
// 		log.Fatalf("Error loading .env file")
// 	}

// }
type Student struct {
	Name    string  `parquet:"name=name, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	Age     int32   `parquet:"name=age, type=INT32, encoding=PLAIN"`
	Id      int64   `parquet:"name=id, type=INT64"`
	Weight  float32 `parquet:"name=weight, type=FLOAT"`
	Sex     bool    `parquet:"name=sex, type=BOOLEAN"`
	Day     int32   `parquet:"name=day, type=INT32, convertedtype=DATE"`
	Ignored int32   //without parquet tag and won't write
}

func main() {
	// chanIn := make(chan []arrow.Record, 20)
	// var wg = &sync.WaitGroup{}
	// // wg.Add(1)
	// // go func() {
	// // 	defer wg.Done()
	// // 	parquet.CsvToArrowChannel("hotels_data.csv", chanIn)
	// // }()

	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	parquet.JsonFileToArrowChannel("hotels_5.json", chanIn)
	// }()

	// wg.Wait()
	// close(chanIn)

	// for record := range chanIn {
	// 	parquet.ArrowToParquet(record)
	// }

	// obj := `[
	// 	{
	// 		"checkin_date": "2024-06-01",
	// 		"checkout_date": "2024-06-07",
	// 		"hotel": {
	// 			"name": "Grand Hotel Luxor",
	// 			"city": "Metropolis",
	// 			"review": 4.3
	// 		}
	// 	},
	// 	{
	// 		"checkin_date": "2024-07-15",
	// 		"checkout_date": "2024-07-22",
	// 		"hotel": {
	// 			"name": "Sunset Beach Resort",
	// 			"city": "Seaville",
	// 			"review": 4.5
	// 		}
	// 	},
	// 	{
	// 		"checkin_date": "2024-08-05",
	// 		"checkout_date": "2024-08-12",
	// 		"hotel": {
	// 			"name": "Mountain View Lodge",
	// 			"city": "Peaksville",
	// 			"review": 4.2
	// 		}
	// 	},
	// 	{
	// 		"checkin_date": "2024-09-10",
	// 		"checkout_date": "2024-09-17",
	// 		"hotel": {
	// 			"name": "Royal Oasis Palace",
	// 			"city": "Kingstown",
	// 			"review": 4.7
	// 		}
	// 	},
	// 	{
	// 		"checkin_date": "2024-10-20",
	// 		"checkout_date": "2024-10-27",
	// 		"hotel": {
	// 			"name": "Golden Sands Retreat",
	// 			"city": "Sunnydale",
	// 			"review": 4.0
	// 		}
	// 	}
	// ]

	// `
	// var Booking []hotel.Booking
	// json.Unmarshal([]byte(obj), &Booking)
	// parquet.StructToParquet("parquetgo", Booking)

	obj := `[
		{
			"name": "Grand Hotel Luxor",
			"city": "Metropolis",
			"review": 100
		}
	]`
	var Hotel []hotel.Hotel
	json.Unmarshal([]byte(obj), &Hotel)
	arrow := parquet.HotelBookingStructToJsonAndArrow(Hotel[0], schema.Fields)
	parquet.ArrowToParquetHotel(arrow, schema.Fields)
	///parquetgo
	// obj := `[
	// 	{
	// 		"id": 1,
	// 		"name": "Grand Hotel Luxor",
	// 		"city": "Metropolis",
	// 		"review": 4.3
	// 	},
	// 	{
	// 		"id": 2,
	// 		"name": "Sunset Beach Resort",
	// 		"city": "Seaville",
	// 		"review": 4.5
	// 	},
	// 	{
	// 		"id": 3,
	// 		"name": "Mountain View Lodge",
	// 		"city": "Peaksville",
	// 		"review": 4.2
	// 	},
	// 	{
	// 		"id": 4,
	// 		"name": "Royal Oasis Palace",
	// 		"city": "Kingstown",
	// 		"review": 4.7
	// 	},
	// 	{
	// 		"id": 5,
	// 		"name": "Golden Sands Retreat",
	// 		"city": "Sunnydale",
	// 		"review": 4.0
	// 	}
	// ]
	// `
	// var Hotel []hotel.Hotel
	// json.Unmarshal([]byte(obj), &Hotel)
	// parquet.WriteParquetFromStruct(Hotel)
	// parquet.ReadParquet()

}
