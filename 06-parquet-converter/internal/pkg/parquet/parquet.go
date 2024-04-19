package parquet

import (
	"fmt"
	"os"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/apache/arrow/go/v16/parquet"
	"github.com/apache/arrow/go/v16/parquet/compress"
	"github.com/apache/arrow/go/v16/parquet/pqarrow"
	"parquet.example/internal/pkg/schema"
	"parquet.example/internal/pkg/utils"
)

func ArrowToParquet(records <-chan []arrow.Record) {

	randomHex, _ := utils.NewRandomSuffix()
	fileName := fmt.Sprintf(utils.GetParquetFilePath()+"hotels_metadata_%s.parquet", randomHex)
	pqout, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	wr, err := pqarrow.NewFileWriter(schema.GetRecordSchema(), pqout,
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

	for recArr := range records {
		for _, rec := range recArr {
			wr.Write(rec)
			rec.Release()
		}

	}
}
