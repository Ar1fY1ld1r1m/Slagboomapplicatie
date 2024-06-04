package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() error {
	dbUser := "root"
	dbPassword := "P@ssword2024"
	dbName := "registraties"
	dbHost := "127.0.0.1"
	dbPort := "3306"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	fmt.Println("DSN:", dsn)

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("Database is ontoegankelijk: %w", err)
	}

	err = DB.Ping()
	if err != nil {
		return fmt.Errorf("Database is ontoegankelijk: %w", err)
	}

	fmt.Println("Verbinding gemaakt met de database!")
	return nil
}

func CheckKentekenInDatabase(kenteken string) (bool, string, string, string, error) {
	var found bool
	var voornaam, achternaam, telefoonnummer string
	query := "SELECT EXISTS (SELECT 1 FROM kentekenhouders WHERE kenteken = ?)"
	err := DB.QueryRow(query, kenteken).Scan(&found)
	if err != nil {
		return false, "", "", "", fmt.Errorf("Fout bij het ophalen van het kenteken: %w", err)
	}

	if found {
		query = "SELECT voornaam, achternaam, telefoonnummer FROM kentekenhouders WHERE kenteken = ?"
		err = DB.QueryRow(query, kenteken).Scan(&voornaam, &achternaam, &telefoonnummer)
		if err != nil {
			return false, "", "", "", fmt.Errorf("Fout bij het ophalen van gegevens: %w", err)
		}
	}

	return found, voornaam, achternaam, telefoonnummer, nil
}
