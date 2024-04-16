package parquet

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/apache/arrow/go/v16/arrow/array"
	"github.com/apache/arrow/go/v16/arrow/memory"
	"github.com/apache/arrow/go/v16/parquet"
	"github.com/apache/arrow/go/v16/parquet/pqarrow"
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
