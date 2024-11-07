package list

import (
	"github.com/jmoiron/sqlx"
	"log"
	"os"
	"strings"
)

type InitializeTables struct{}

func (m *InitializeTables) GetName() string {
	return "InitializeTables"
}

func (m *InitializeTables) Up(con *sqlx.DB) {
	file, err := os.ReadFile(migrationFileLocation() + "initialize_tables.sql")
	if err != nil {
		log.Fatalf("Unexpected error reading file: %v", err.Error())
	}

	// Explode the queries into individual ones
	queries := strings.Split(string(file), ";")

	for _, query := range queries {
		_, err = con.Exec(query)
		if err != nil {
			log.Printf("Error executing a query in initialize_tables: %v\nQuery:%v", err.Error(), query)
		}
	}

	_, err = con.Exec(string(file))
	if err != nil {
		log.Printf("Did not execute initialize tables migration: %v", err.Error())
	}
}

func (m *InitializeTables) Down(con *sqlx.DB) {
	con.MustExec(`DROP DATABASE DB_CMPPlus_MS;`)
}
