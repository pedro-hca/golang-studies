package parquet

import (
	// "github.com/apache/arrow/go/arrow"
	// "github.com/apache/arrow/go/arrow/array"
	// "github.com/apache/arrow/go/arrow/memory"
	// "github.com/apache/arrow/go/arrow/parquet/pqarrow"
	"fmt"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/apache/arrow/go/v16/arrow/array"
	"github.com/apache/arrow/go/v16/arrow/memory"
)

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
