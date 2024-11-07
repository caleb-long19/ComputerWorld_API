package list

import (
	"github.com/jmoiron/sqlx"
)

type CreateUser struct{}

func (m *CreateUser) GetName() string {
	return "CreateUser"
}

func (m *CreateUser) Up(con *sqlx.DB) {
	con.MustExec(`CREATE TABLE user(
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

func (m *CreateUser) Down(con *sqlx.DB) {
	con.MustExec(`DROP TABLE user;`)
}
