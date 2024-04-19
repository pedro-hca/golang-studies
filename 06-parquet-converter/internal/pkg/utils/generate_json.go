package utils

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

type Hotel struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	City   string  `json:"city"`
	Review float64 `json:"review"`
}

func GenerateJson() {
	file, err := os.Create(GetJsonFilePath() + "hotels_100000.json")
	if err != nil {
		fmt.Println("Error while creating json file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	// Gerar 10 mil registros
	for i := 1; i <= 100000; i++ {
		hotel := Hotel{
			ID:     i,
			Name:   fmt.Sprintf("Hotel %d", i),
			City:   fmt.Sprintf("City %d", rand.Intn(100)),
			Review: rand.Float64() * 5, // Gerar uma avaliação aleatória entre 0 e 5
		}

		if err := encoder.Encode(hotel); err != nil {
			fmt.Println("Error while enconding hotel:", err)
			return
		}
	}

	fmt.Println(".json is ready!")
}

func GenerateRecords(numRecords int) []Hotel {
	var records []Hotel
	for i := 1; i <= numRecords; i++ {
		record := Hotel{
			ID:     i,
			Name:   "Hotel" + strconv.Itoa(i),
			City:   "City" + strconv.Itoa(i),
			Review: float64(i) * 0.1, // Geração de revisão fictícia
		}
		records = append(records, record)
	}
	return records
}

func WriteCSV(filename string, records []Hotel) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Escrever cabeçalho
	header := []string{"id", "name", "city", "review"}
	if err := writer.Write(header); err != nil {
		return err
	}

	// Escrever registros
	for _, record := range records {
		row := []string{
			strconv.Itoa(record.ID),
			record.Name,
			record.City,
			strconv.FormatFloat(record.Review, 'f', 1, 64),
		}
		if err := writer.Write(row); err != nil {
			return err
		}
	}

	return nil
}
