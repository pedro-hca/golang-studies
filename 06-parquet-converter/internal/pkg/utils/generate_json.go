package utils

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
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
