package parquet

import (
	"fmt"
	"log"

	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/parquet"
	"github.com/xitongsys/parquet-go/reader"
	"github.com/xitongsys/parquet-go/writer"
	"parquet.example/internal/pkg/hotel"
)

func WriteParquetFromStruct(Hotels []hotel.Hotel) {
	var err error
	fw, err := local.NewLocalFileWriter("flat.parquet")
	if err != nil {
		log.Println("Can't create local file", err)
		return
	}

	// write
	pw, err := writer.NewParquetWriter(fw, new(hotel.Hotel), 4)
	if err != nil {
		log.Println("Can't create parquet writer", err)
		return
	}

	pw.RowGroupSize = 128 * 1024 * 1024 //128M
	pw.PageSize = 8 * 1024              //8K
	pw.CompressionType = parquet.CompressionCodec_SNAPPY
	for _, hotel := range Hotels {

		err = pw.Write(hotel)
		if err != nil {
			log.Println("Write error", err)
		}
	}
	err = pw.WriteStop()
	if err != nil {
		log.Println("WriteStop error", err)
		return
	}
	log.Println("Write Finished")
	fw.Close()
}

func ReadParquet() {
	// /read
	fr, err := local.NewLocalFileReader("flat.parquet")
	if err != nil {
		log.Println("Can't open file")
		return
	}

	pr, err := reader.NewParquetReader(fr, new(hotel.Hotel), 4)
	if err != nil {
		log.Println("Can't create parquet reader", err)
		return
	}
	num := int(pr.GetNumRows())
	hotels := make([]hotel.Hotel, num)
	for i := 0; i < num; i++ {

		hotelsRead := make([]hotel.Hotel, 1)
		err = pr.Read(&hotelsRead)
		if err != nil {
			log.Println("Read error", err)
		}
		log.Println(hotelsRead)
		hotels[i] = hotelsRead[0]
	}
	fmt.Println(hotels)

	pr.ReadStop()
	fr.Close()
}
