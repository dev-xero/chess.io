package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/dev-xero/chess.io/internal/config"
	"github.com/dev-xero/chess.io/internal/database"
	"github.com/dev-xero/chess.io/internal/models"
)

func main() {
	file, err := os.Open("cmd/seed/iso.csv")
	if err != nil {
		log.Fatal("Failed to open 'iso.csv' file: ", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Error reading 'iso.csv' file: ", err)
	}

	var countries []models.Country
	for _, record := range records[1:] {
		countries = append(countries, models.Country{
			Name: record[0],
			Code: record[2],
		})
	}

	cfg := config.Load()
	db := database.Connect(cfg)

	result := db.Create(&countries)
	if result.Error != nil {
		log.Fatal("Failed to seed countries: ", result.Error)
	}

	log.Println("\nSeeded countries table successfully")
}
