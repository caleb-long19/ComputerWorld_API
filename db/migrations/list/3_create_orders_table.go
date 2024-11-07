package list

import (
	"github.com/jmoiron/sqlx"
)

type CreateOrders struct{}

func (m *CreateOrders) GetName() string {
	return "CreateOrders"
}

func (m *CreateOrders) Up(con *sqlx.DB) {
	con.MustExec(`CREATE TABLE orders(
		uid VARCHAR(36) NOT NULL PRIMARY KEY,
		order_ref VARCHAR(500) NOT NULL,
		order_amount int NOT NULL,
		product_uid VARCHAR(36) NOT NULL,
		order_price float NOT NULL,
		created_at DATETIME default current_timestamp,
		updated_at DATETIME default current_timestamp,
		deleted_at DATETIME DEFAULT NULL,
		)
	`)
}

func (m *CreateOrders) Down(con *sqlx.DB) {
	con.MustExec(`DROP TABLE products;`)
}
