package parquet

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/apache/arrow/go/v16/arrow/array"
	"github.com/apache/arrow/go/v16/arrow/memory"
	"github.com/apache/arrow/go/v16/parquet"
	"github.com/apache/arrow/go/v16/parquet/pqarrow"
	"parquet.example/internal/pkg/hotel"
	"parquet.example/internal/pkg/utils"
)

// just testing int build
func BuildInt64() {
	bldr := array.NewInt64Builder(memory.DefaultAllocator)
	defer bldr.Release()

	bldr.Append(25)
	bldr.AppendNull()
	// nil bool slice mean ALL valid
	bldr.AppendValues([]int64{1, 2, 3, 4}, nil)
	// otherwise bool slice indicates nulls with false
	bldr.AppendValues([]int64{5, 0, 6, 7},
		[]bool{true, false, true, true})

	arr := bldr.NewArray()
	defer arr.Release()
	fmt.Println(arr)
	// Output: [25 (null) 1 2 3 4 5 (null) 6 7]
}

// just testing build structs
func BuildStruct(names []string, cities []string, reviews []float32) {
	// Schema
	hotelType := arrow.StructOf(
		arrow.Field{Name: "name", Type: arrow.BinaryTypes.String},
		arrow.Field{Name: "city", Type: arrow.BinaryTypes.String},
		arrow.Field{Name: "review", Type: arrow.PrimitiveTypes.Float32},
	)

	bldr := array.NewStructBuilder(memory.DefaultAllocator, hotelType) // DefaultAllocator means go memory alocator, like a make() function
	defer bldr.Release()
	// notice we don't have to separetely realse these
	// they are owned by the struct builder!
	namesBldr := bldr.FieldBuilder(0).(*array.StringBuilder)
	citiesBldr := bldr.FieldBuilder(1).(*array.StringBuilder)
	reviewsBldr := bldr.FieldBuilder(2).(*array.Float32Builder)

	for i := range names {
		bldr.Append(true) // <-- Valid Struct
		namesBldr.Append(names[i])
		citiesBldr.Append(cities[i])
		reviewsBldr.Append(reviews[i])
	}

	bldr.Append(false) // == bldr.AppendNull()

	arr := bldr.NewStructArray()
	defer arr.Release() // new array! need to release!

	fmt.Println(arr)

}

func CsvJsonToParquetGoroutines() {
	/// implementar essa go routine
	var wg sync.WaitGroup
	const nworkers = 10
	wg.Add(nworkers)
	for i := 0; i < nworkers; i++ {
		go func() {
			defer wg.Done()
			// put the record manipulation here
		}()
	}

	go func() {
		wg.Wait()
		// close(ch2)
	}()

	/// need to implement
}

func ReadParquetFile() {
	f, err := os.Open(utils.GetParquetFilePath() + "hotels_metadata_79256469048fc715.parquet")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	tbl, err := pqarrow.ReadTable(context.Background(), f, parquet.NewReaderProperties(memory.DefaultAllocator),
		pqarrow.ArrowReadProperties{}, memory.DefaultAllocator)
	if err != nil {
		panic(err)
	}

	s := tbl.Schema()
	fmt.Println(s)
	fmt.Println("------")

	fmt.Println("the count of table columns=", tbl.NumCols())
	fmt.Println("the count of table rows=", tbl.NumRows())
	fmt.Println("------")

	// for i := 0; i < int(tbl.NumCols()); i++ {
	// 	col := tbl.Column(i)
	// 	fmt.Printf("arrays in column(%s):\n", col.Name())
	// 	chunk := col.Data()
	// 	for _, arr := range chunk.Chunks() {
	// 		fmt.Println(arr)
	// 	}
	// 	fmt.Println("------")
	// }

}

func HotelBookingStructToJsonAndArrow(flatHotelBookingData []map[string]interface{}, schemaFields []arrow.Field) []arrow.Record {
	var recordArray []arrow.Record

	schemaRecord := arrow.NewSchema(schemaFields, nil)
	schemaStruct := arrow.StructOf(schemaFields...)

	bldr := array.NewStructBuilder(memory.DefaultAllocator, schemaStruct)
	defer bldr.Release()

	flatByte, _ := json.Marshal(flatHotelBookingData)
	jsonArray := "[" + string(flatByte) + "]"
	err := bldr.UnmarshalJSON([]byte(jsonArray))
	fmt.Println(err)

	structArray := bldr.NewStructArray()
	defer structArray.Release()
	fmt.Println(structArray)

	recordArray = append(recordArray, array.RecordFromStructArray(structArray, schemaRecord))
	fmt.Println(recordArray)
	return recordArray
}
func HotelBookingStructToArrow(flatHotelBookingData hotel.Hotel, schemaFields []arrow.Field) []arrow.Record {
	var recordArray []arrow.Record

	schemaRecord := arrow.NewSchema(schemaFields, nil)
	schemaStruct := arrow.StructOf(schemaFields...)

	bldr := array.NewStructBuilder(memory.DefaultAllocator, schemaStruct)
	defer bldr.Release()

	travelerIDBldr := bldr.FieldBuilder(0)
	idBldr := bldr.FieldBuilder(1)
	bookingEndTimeLocalBldr := bldr.FieldBuilder(2)

	bldr.Append(true)
	travelerIDBldr.(*array.StringBuilder).Append(flatHotelBookingData.Name)
	idBldr.(*array.StringBuilder).Append(flatHotelBookingData.City)

	// Verifica e converte o Review para float64 se necessário
	// reviewValue := flatHotelBookingData.Review
	// if reviewValue == float64(int(reviewValue)) { // Verificação redundante
	// 	reviewValue = float64(int(reviewValue))
	// }
	// reviewValue = reviewValue + 0.1
	bookingEndTimeLocalBldr.(*array.Float64Builder).Append(flatHotelBookingData.Review)

	// bookingEndTimeLocalBldr.(*array.Float64Builder).Append(flatHotelBookingData.Review)

	structArray := bldr.NewStructArray()
	defer structArray.Release()
	fmt.Println(structArray)

	recordArray = append(recordArray, array.RecordFromStructArray(structArray, schemaRecord))
	fmt.Println(recordArray)
	return recordArray
}

func ArrowToParquetHotel(records []arrow.Record, schemaFields []arrow.Field) {
	pqout, err := os.Create("FlatHotelBookingData.parquet")
	if err != nil {
		panic(err)
	}
	defer pqout.Close()
	schemaRecord := arrow.NewSchema(schemaFields, nil)
	wr, err := pqarrow.NewFileWriter(schemaRecord, pqout,
		parquet.NewWriterProperties(), pqarrow.DefaultWriterProps())
	if err != nil {
		panic(err)
	}
	defer wr.Close()

	for _, rec := range records {
		fmt.Println(rec)
		wr.Write(rec)
		rec.Release()

	}
}
