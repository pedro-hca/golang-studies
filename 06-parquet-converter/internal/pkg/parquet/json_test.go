package parquet_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"parquet.example/internal/pkg/parquet"
	"parquet.example/internal/pkg/utils"
)

func BenchmarkJsonFileToParquet(b *testing.B) {
	outputDir := utils.GetTestFilePath()
	err := setup(outputDir)
	if err != nil {
		b.Fatalf("Setup error: %v", err)
	}

	filePath := "hotels.json" // Ensure this file exists with representative data

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := parquet.JsonFileToParquet(filePath)
		if err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()

	err = teardown(outputDir)
	if err != nil {
		b.Fatalf("Teardown error: %v", err)
	}
}

func setup(outputDir string) error {
	dir := filepath.Dir(outputDir)
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return fmt.Errorf("output file path does not exist: %w", err)
	}
	return nil
}

func teardown(outputDir string) error {
	return filepath.Walk(outputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(path) == ".parquet" {
			return os.Remove(path)
		}
		return nil
	})
}
