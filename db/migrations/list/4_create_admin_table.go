package list

import (
	"github.com/jmoiron/sqlx"
)

type CreateAdmin struct{}

func (m *CreateAdmin) GetName() string {
	return "CreateAdmin"
}

func (m *CreateAdmin) Up(con *sqlx.DB) {
	con.MustExec(`CREATE TABLE admin(
		uid VARCHAR(36) NOT NULL PRIMARY KEY,
		email VARCHAR(255) NOT NULL,
		name VARCHAR(255) NOT NULL,
		password VARCHAR(255) NOT NULL,
		created_at DATETIME default current_timestamp,
		updated_at DATETIME default current_timestamp,
		deleted_at DATETIME DEFAULT NULL,
		)
	`)
}

func (m *CreateAdmin) Down(con *sqlx.DB) {
	con.MustExec(`DROP TABLE admin;`)
}
