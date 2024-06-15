package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	DBUsername string `json:"db_username"`
	DBPassword string `json:"db_password"`
	DBHost     string `json:"db_host"`
	DBPort     string `json:"db_port"`
	DBName     string `json:"db_name"`
}

var DB *sql.DB

func Connect() error {
	err := connectDatabase("cloudconfig.json")
	if err != nil {
		err = connectDatabase("localconfig.json")
	}
	return err
}

func connectDatabase(configFile string) error {
	file, err := os.Open(configFile)
	if err != nil {
		return fmt.Errorf("error opening config file: %w", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var config Config
	err = decoder.Decode(&config)
	if err != nil {
		return fmt.Errorf("error decoding config JSON: %w", err)
	}

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DBUsername, config.DBPassword, config.DBHost, config.DBPort, config.DBName)
	DB, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		return fmt.Errorf("database is ontoegankelijk: %w", err)
	}

	err = DB.Ping()
	if err != nil {
		return fmt.Errorf("database is ontoegankelijk: %w", err)
	}

	return nil
}

func CheckKentekenInDatabase(kenteken string) (bool, string, string, string, error) {
	var found bool
	var voornaam, achternaam, telefoonnummer string
	query := "SELECT EXISTS (SELECT 1 FROM reservations WHERE license_plate = ?)"
	err := DB.QueryRow(query, kenteken).Scan(&found)
	if err != nil {
		return false, "", "", "", fmt.Errorf("fout bij het ophalen van het kenteken: %w", err)
	}

	if found {
		query = "SELECT name, surname, phone FROM reservations WHERE license_plate = ?"
		err = DB.QueryRow(query, kenteken).Scan(&voornaam, &achternaam, &telefoonnummer)
		if err != nil {
			return false, "", "", "", fmt.Errorf("fout bij het ophalen van gegevens: %w", err)
		}
	}

	return found, voornaam, achternaam, telefoonnummer, nil
}
