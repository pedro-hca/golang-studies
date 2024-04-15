package parquet

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"os"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/apache/arrow/go/v16/arrow/array"
	"github.com/apache/arrow/go/v16/arrow/memory"
	"github.com/apache/arrow/go/v16/parquet"
	"github.com/apache/arrow/go/v16/parquet/compress"
	"github.com/apache/arrow/go/v16/parquet/pqarrow"
)

func JsonFileToParquet() error {
	var recordArray []arrow.Record

	jsonFile, err := os.Open("../csv/hotels.json")
	if err != nil {
		return fmt.Errorf("Error while opening .json file: %v", err)
	}
	defer jsonFile.Close()

	jsonBytes, err := io.ReadAll(jsonFile)
	if err != nil {
		return fmt.Errorf("Error while reading .json file: %v", err)
	}
	// Schema Record
	schemaRecord := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int32},
			{Name: "name", Type: arrow.BinaryTypes.String},
			{Name: "city", Type: arrow.BinaryTypes.String},
			{Name: "review", Type: arrow.PrimitiveTypes.Float64},
		}, nil,
	)
	// Schema Struct
	schemaStruct := arrow.StructOf(
		arrow.Field{Name: "id", Type: arrow.PrimitiveTypes.Int32},
		arrow.Field{Name: "name", Type: arrow.BinaryTypes.String},
		arrow.Field{Name: "city", Type: arrow.BinaryTypes.String},
		arrow.Field{Name: "review", Type: arrow.PrimitiveTypes.Float64},
	)
	rb := array.NewStructBuilder(memory.DefaultAllocator, schemaStruct)
	defer rb.Release() // Ensure the builder releases its resources

	rb.UnmarshalJSON(jsonBytes)
	structArray := rb.NewStructArray()
	defer structArray.Release() // Match retain with release

	recordArray = append(recordArray, array.RecordFromStructArray(structArray, schemaRecord))

	pqout, err := os.Create("hotels_json_file_metadata.parquet")
	if err != nil {
		panic(err)
	}
	defer pqout.Close()

	wr, err := pqarrow.NewFileWriter(schemaRecord, pqout,
		parquet.NewWriterProperties(
			parquet.WithCompression(compress.Codecs.Snappy),
			parquet.WithCompressionFor("review", compress.Codecs.Zstd),
			parquet.WithDictionaryDefault(false),
			parquet.WithDictionaryFor("city", true),
			parquet.WithEncodingFor("id", parquet.Encodings.DeltaBinaryPacked),
			parquet.WithDataPageVersion(parquet.DataPageV2),
			parquet.WithVersion(parquet.V2_LATEST),
		), pqarrow.DefaultWriterProps())
	if err != nil {
		panic(err)
	}
	defer wr.Close()
	for _, rec := range recordArray {
		err := wr.Write(rec)
		if err != nil {
			panic(err)
		}
		rec.Release()
	}

	return nil

}
func JsonToParquet() {
	jsonObj := `[
		{
			"id": 1,
			"name": "Grand Hotel Luxor",
			"city": "Metropolis",
			"review": 4.3
		},
		{
			"id": 2,
			"name": "Sunset Beach Resort",
			"city": "Seaville",
			"review": 4.5
		},
		{
			"id": 3,
			"name": "Mountain View Lodge",
			"city": "Peaksville",
			"review": 4.2
		},
		{
			"id": 4,
			"name": "Royal Oasis Palace",
			"city": "Kingstown",
			"review": 4.7
		},
		{
			"id": 5,
			"name": "Golden Sands Retreat",
			"city": "Sunnydale",
			"review": 4.0
		}
	]
	`
	var recordArray []arrow.Record

	// Schema Record
	schemaRecord := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int32},
			{Name: "name", Type: arrow.BinaryTypes.String},
			{Name: "city", Type: arrow.BinaryTypes.String},
			{Name: "review", Type: arrow.PrimitiveTypes.Float64},
		}, nil,
	)
	// Schema Struct
	schemaStruct := arrow.StructOf(
		arrow.Field{Name: "id", Type: arrow.PrimitiveTypes.Int32},
		arrow.Field{Name: "name", Type: arrow.BinaryTypes.String},
		arrow.Field{Name: "city", Type: arrow.BinaryTypes.String},
		arrow.Field{Name: "review", Type: arrow.PrimitiveTypes.Float64},
	)
	rb := array.NewStructBuilder(memory.DefaultAllocator, schemaStruct)

	rb.UnmarshalJSON([]byte(jsonObj))
	structArray := rb.NewStructArray()
	recordArray = append(recordArray, array.RecordFromStructArray(structArray, schemaRecord))

	pqout, err := os.Create("hotels_json_obj_metadata.parquet")
	if err != nil {
		panic(err)
	}

	wr, err := pqarrow.NewFileWriter(schemaRecord, pqout,
		parquet.NewWriterProperties(
			parquet.WithCompression(compress.Codecs.Snappy),
			parquet.WithCompressionFor("review", compress.Codecs.Zstd),
			parquet.WithDictionaryDefault(false),
			parquet.WithDictionaryFor("city", true),
			parquet.WithEncodingFor("id", parquet.Encodings.DeltaBinaryPacked),
			parquet.WithDataPageVersion(parquet.DataPageV2),
			parquet.WithVersion(parquet.V2_LATEST),
		), pqarrow.DefaultWriterProps())
	if err != nil {
		panic(err)
	}
	defer wr.Close()
	for _, rec := range recordArray {
		err := wr.Write(rec)
		if err != nil {
			panic(err)
		}
		rec.Release()
	}

}

func JsonToParquetGoroutines() {
	jsonObj := `[
		{
			"id": 1,
			"name": "Grand Hotel Luxor",
			"city": "Metropolis",
			"review": 4.3
		},
		{
			"id": 2,
			"name": "Sunset Beach Resort",
			"city": "Seaville",
			"review": 4.5
		},
		{
			"id": 3,
			"name": "Mountain View Lodge",
			"city": "Peaksville",
			"review": 4.2
		},
		{
			"id": 4,
			"name": "Royal Oasis Palace",
			"city": "Kingstown",
			"review": 4.7
		},
		{
			"id": 5,
			"name": "Golden Sands Retreat",
			"city": "Sunnydale",
			"review": 4.0
		}
	]
	`
	var recordArray []arrow.Record

	// Schema Record
	schemaRecord := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int32},
			{Name: "name", Type: arrow.BinaryTypes.String},
			{Name: "city", Type: arrow.BinaryTypes.String},
			{Name: "review", Type: arrow.PrimitiveTypes.Float64},
		}, nil,
	)
	// Schema Struct
	schemaStruct := arrow.StructOf(
		arrow.Field{Name: "id", Type: arrow.PrimitiveTypes.Int32},
		arrow.Field{Name: "name", Type: arrow.BinaryTypes.String},
		arrow.Field{Name: "city", Type: arrow.BinaryTypes.String},
		arrow.Field{Name: "review", Type: arrow.PrimitiveTypes.Float64},
	)
	rb := array.NewStructBuilder(memory.DefaultAllocator, schemaStruct)
	defer rb.Release()

	rb.UnmarshalJSON([]byte(jsonObj))
	structArray := rb.NewStructArray()
	// structArray.Retain()
	defer structArray.Release()

	recordChan := make(chan []arrow.Record)
	doneChan := make(chan bool)

	go func() {
		// for i := 0; i < structArray.Len(); i++ {
		recordArray := append(recordArray, array.RecordFromStructArray(structArray, schemaRecord))
		recordChan <- recordArray
		// }

		defer close(recordChan)
	}()
	go func() {

		for recordArrayChan := range recordChan {

			// Generate 8 random bytes
			randomBytes := make([]byte, 8)
			_, err := rand.Read(randomBytes)
			if err != nil {
				fmt.Errorf("Error while generating random suffix name", err)
				return
			}
			// Convert bytes to hexadecimal string
			randomHex := hex.EncodeToString(randomBytes)
			fileName := fmt.Sprintf("hotels_metadata_%s.parquet", randomHex)

			pqout, err := os.Create(fileName)
			if err != nil {
				panic(err)
			}

			wr, err := pqarrow.NewFileWriter(schemaRecord, pqout,
				parquet.NewWriterProperties(
					parquet.WithCompression(compress.Codecs.Snappy),
					parquet.WithCompressionFor("review", compress.Codecs.Zstd),
					parquet.WithDictionaryDefault(false),
					parquet.WithDictionaryFor("city", true),
					parquet.WithEncodingFor("id", parquet.Encodings.DeltaBinaryPacked),
					parquet.WithDataPageVersion(parquet.DataPageV2),
					parquet.WithVersion(parquet.V2_LATEST),
				), pqarrow.DefaultWriterProps())
			if err != nil {
				panic(err)
			}
			defer wr.Close()

			for _, rec := range recordArrayChan {

				err := wr.Write(rec)
				if err != nil {
					panic(err)
				}
				rec.Release()
			}
		}
		doneChan <- true
	}()
	<-doneChan

}
