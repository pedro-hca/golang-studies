package parquet

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/apache/arrow/go/v16/arrow/array"
	"github.com/apache/arrow/go/v16/arrow/csv"
	"github.com/apache/arrow/go/v16/arrow/memory"
	"github.com/apache/arrow/go/v16/parquet"
	"github.com/apache/arrow/go/v16/parquet/compress"
	"github.com/apache/arrow/go/v16/parquet/pqarrow"
	"parquet.example/internal/pkg/schema"
	"parquet.example/internal/pkg/utils"
)

func CsvToParquet(filePath string) {
	var recordArr []arrow.Record
	ch := make(chan []arrow.Record, 20)
	schemaRecord := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "name", Type: arrow.BinaryTypes.String},
			{Name: "city", Type: arrow.BinaryTypes.String},
			{Name: "review", Type: arrow.PrimitiveTypes.Float64},
		}, nil,
	)
	go func() {
		//close the channel when done to signal
		defer close(ch)

		file, err := os.Open(utils.GetCsvFilePath() + filePath)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		// infer the types and schema from the header line
		// and first line of data
		rdr := csv.NewInferringReader(file, csv.WithChunk(-1),
			// strings can be null, and these are the values
			// to consider as "null"
			csv.WithNullReader(true, "", "null", "[]"),
			csv.WithHeader(true))

		for rdr.Next() {
			rec := rdr.Record()
			structArray := array.RecordToStructArray(rec)
			fmt.Println(rec)
			recordArr = append(recordArr, array.RecordFromStructArray(structArray, schemaRecord))
			ch <- recordArr
		}

		if rdr.Err() != nil {
			panic(rdr.Err())
		}
	}()

	randomHex, err := utils.NewRandomSuffix()
	if err != nil {
		innerErr := errors.Unwrap(err)
		log.Fatalf("Iternal error: %v", innerErr)

	}
	fileName := fmt.Sprintf(utils.GetParquetFilePath()+"hotels_metadata_%s.parquet", randomHex)
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

	for recArr := range ch {
		for _, rec := range recArr {
			wr.Write(rec)
			rec.Release()
		}

	}
}
func CsvToArrow(fileName string) chan []arrow.Record {
	var recordArr []arrow.Record
	out := make(chan []arrow.Record)
	go func() {
		//close the channel when done to signal
		defer close(out)

		file, err := os.Open(utils.GetCsvFilePath() + fileName)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		// infer the types and schema from the header line
		// and first line of data
		rdr := csv.NewInferringReader(file, csv.WithChunk(-1),
			// strings can be null, and these are the values
			// to consider as "null"
			csv.WithNullReader(true, "", "null", "[]"),
			csv.WithHeader(true))

		for rdr.Next() {
			fmt.Println(rdr)
			rec := rdr.Record()
			structArray := array.RecordToStructArray(rec)
			fmt.Println(rec)
			recordArr = append(recordArr, array.RecordFromStructArray(structArray, schema.GetRecordSchema()))
			out <- recordArr
		}

		if rdr.Err() != nil {
			panic(rdr.Err())
		}
	}()
	return out

}

func CsvFileToParquet() {

	ch := make(chan arrow.Record, 20)
	ch2 := make(chan arrow.Record, 1)
	ch3 := make(chan arrow.Record, 1)
	go func() {
		//close the channel when done to signal
		// future pipeline steps
		defer close(ch)
		file, err := os.Open(utils.GetCsvFilePath() + "film.csv")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		// infer the types and schema from the header line
		// and first line of data

		rdr := csv.NewInferringReader(file, csv.WithChunk(-1),
			// strings can be null, and these are the values
			// to consider as "null"
			csv.WithNullReader(true, "", "null", "[]"),
			csv.WithHeader(true))

		for rdr.Next() {
			fmt.Println(rdr)
			rec := rdr.Record()
			fmt.Println(rec)
			rec.Retain()
			ch <- rec
		}

		if rdr.Err() != nil {
			panic(rdr.Err())
		}
	}()

	// we need to know the fields we're expecting in this JSON string
	// harcoded
	bldr := array.NewListBuilder(memory.DefaultAllocator, arrow.StructOf(
		arrow.Field{Name: "id", Type: arrow.PrimitiveTypes.Int64},
		arrow.Field{Name: "name", Type: arrow.BinaryTypes.String},
	))
	fmt.Println(bldr)

	defer bldr.Release()
	fmt.Println(bldr)

	var outSchema *arrow.Schema
	for rec := range ch {
		genresCol := rec.Column(0).(*array.String)

		bldr.Reserve(int(rec.NumRows()))
		for i := 0; i < genresCol.Len(); i++ {
			if genresCol.IsNull(i) {
				bldr.AppendNull()
				continue
			}

			re := regexp.MustCompile(`'([^']*)'`)
			vals := re.ReplaceAllString(genresCol.Value(i), `"$1"`)

			if !json.Valid([]byte(vals)) {
				panic("invalid JSON")
			}
			err := bldr.UnmarshalJSON([]byte("[" + vals + "]"))
			if err != nil {
				panic(err)
			}
		}

		cols := rec.Columns()
		// modify the slice of arrays
		// new record doesn't copy the colums!
		cols[0].Release()
		genreCol := bldr.NewArray()
		cols[0] = genreCol

		// if we don't know the entire schema beforehand, we can just copy the existing
		// schema and replace the field for the column we're altering
		if outSchema == nil {
			fieldList := make([]arrow.Field, rec.NumCols())
			copy(fieldList, rec.Schema().Fields())
			fieldList[0].Type = bldr.Type()
			meta := rec.Schema().Metadata()
			outSchema = arrow.NewSchema(fieldList, &meta)
		}

		ch2 <- array.NewRecord(outSchema, cols, rec.NumRows())
		ch3 <- array.NewRecord(outSchema, cols, rec.NumRows())
		rec.Release()
	}
	close(ch2)
	close(ch3)

	randomHex, err := utils.NewRandomSuffix()
	if err != nil {
		innerErr := errors.Unwrap(err)
		log.Fatalf("Iternal error: %v", innerErr)

	}
	fileName := fmt.Sprintf(utils.GetParquetFilePath()+"hotels_metadata_%s.parquet", randomHex)
	pqout, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	firstRec := <-ch2

	wr, err := pqarrow.NewFileWriter(firstRec.Schema(), pqout,
		parquet.NewWriterProperties(
			parquet.WithCompression(compress.Codecs.Snappy),
			parquet.WithCompressionFor("overview", compress.Codecs.Zstd),
			parquet.WithDictionaryDefault(false),
			parquet.WithDictionaryFor("original_language", true),
			parquet.WithDictionaryFor("status", true),
			parquet.WithEncodingFor("id", parquet.Encodings.DeltaBinaryPacked),
			parquet.WithDataPageVersion(parquet.DataPageV2),
			parquet.WithVersion(parquet.V2_LATEST),
		), pqarrow.DefaultWriterProps())
	if err != nil {
		panic(err)
	}
	defer wr.Close()
	firstRec.Release()

	for rec := range ch3 {
		wr.Write(rec)
		rec.Release()
	}
}
