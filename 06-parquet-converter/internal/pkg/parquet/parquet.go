package parquet

import (
	"fmt"
	"os"

	"github.com/parquet-go/parquet-go"
	"parquet.example/internal/pkg/hotel"
)

func StructToParquet(fileName string, hotels []hotel.Hotel) error {
	file := fmt.Sprintf("%s.parquet", fileName)

	err := parquet.WriteFile(file, hotels)
	if err != nil {
		return fmt.Errorf("Error writing file:, %v", err)
	}
	return nil

}

func GenericToParquet(fileName string, hotels []hotel.Hotel) error {
	file := fmt.Sprintf("%s.parquet", fileName)
	fileHotel, err := os.Create(file)
	if err != nil {
		return fmt.Errorf("Error creating file:, %v", err)
	}

	writer := parquet.NewGenericWriter[hotel.Hotel](fileHotel)
	// escrever por cima do arquivo

	_, err = writer.Write(hotels)
	if err != nil {
		return fmt.Errorf("Error writing file:, %v", err)
	}

	err = writer.Close()
	if err != nil {
		return fmt.Errorf("Error writing file:, %v", err)
	}
	return nil

}

// func OverWriteParquet(file os.File, hotels []hotel.Hotel) error {

// }
