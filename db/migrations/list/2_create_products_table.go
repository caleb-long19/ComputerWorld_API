package list

import (
	"github.com/jmoiron/sqlx"
)

type CreateProducts struct{}

func (m *CreateProducts) GetName() string {
	return "CreateProducts"
}

func (m *CreateProducts) Up(con *sqlx.DB) {
	con.MustExec(`CREATE TABLE products(
		uid VARCHAR(36) NOT NULL PRIMARY KEY,
		product_code VARCHAR(500) NOT NULL,
		product_name VARCHAR(500) NOT NULL,
		manufacturer_uid VARCHAR(36) NOT NULL,
		stock int NOT NULL,
		price float NOT NULL,
		created_at DATETIME default current_timestamp,
		updated_at DATETIME default current_timestamp,
		deleted_at DATETIME DEFAULT NULL,
		)
	`)
}

func (m *CreateProducts) Down(con *sqlx.DB) {
	con.MustExec(`DROP TABLE products;`)
}
