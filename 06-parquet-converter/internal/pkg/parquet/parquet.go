package parquet

import (
	"fmt"
	"os"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/apache/arrow/go/v16/parquet"
	"github.com/apache/arrow/go/v16/parquet/pqarrow"
	"parquet.example/internal/pkg/schema"
	"parquet.example/internal/pkg/utils"
)

func ArrowToParquet(records []arrow.Record) {
	randomHex, _ := utils.NewRandomSuffix()
	fileName := fmt.Sprintf(utils.GetParquetFilePath()+"hotels_metadata_%s.parquet", randomHex)
	pqout, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	wr, err := pqarrow.NewFileWriter(schema.GetRecordSchema(), pqout,
		parquet.NewWriterProperties(), pqarrow.DefaultWriterProps())
	if err != nil {
		panic(err)
	}
	defer wr.Close()

	for _, rec := range records {

		wr.Write(rec)
		rec.Release()

	}
}
