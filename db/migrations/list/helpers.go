package list

import (
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

var cachedLoc = ""

func runAndLogError(con *sqlx.DB, query string) {

	_, err := con.Exec(query)

	if err != nil {
		log.Printf("Error running a query: %v\n Query: %v", err, query)
	}
}

// Find the database migration folder
func migrationFileLocation() string {

	// Return the cached loc if it's been set.
	if cachedLoc != "" {
		return cachedLoc
	}

	loc := "db/migrations/sql_files/"
	// find the db folder
	if _, err := os.Stat(loc); err == nil {
		cachedLoc = loc
		return loc
	}

	loc = "../" + loc
	if _, err := os.Stat(loc); err == nil {
		cachedLoc = loc
		return loc
	}

	loc = "../../" + loc
	if _, err := os.Stat(loc); err == nil {
		cachedLoc = loc
		return loc
	}

	log.Panic("Unable to find the database SQL migrations folder")

	return ""
}
