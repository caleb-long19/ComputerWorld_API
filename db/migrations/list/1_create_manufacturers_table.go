package list

import (
	"github.com/jmoiron/sqlx"
)

type CreateManufacturers struct{}

func (m *CreateManufacturers) GetName() string {
	return "CreateManufacturers"
}

func (m *CreateManufacturers) Up(con *sqlx.DB) {
	con.MustExec(`CREATE TABLE manufacturers(
		uid VARCHAR(36) NOT NULL PRIMARY KEY,
		manufacturer_name VARCHAR(500) NOT NULL,
		created_at DATETIME default current_timestamp,
		updated_at DATETIME default current_timestamp,
		deleted_at DATETIME DEFAULT NULL,
		)
	`)
}

func (m *CreateManufacturers) Down(con *sqlx.DB) {
	con.MustExec(`DROP TABLE manufacturers;`)
}
