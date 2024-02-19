package dto

import (
	"database/sql"
	"time"
)

type TransactionSale struct {
	ID              int            `db:"id"`
	SalesmanID      int            `db:"salesman_id"`
	CustomerID      int            `db:"customer_id"`
	TransactionDate time.Time      `db:"transaction_date"`
	Invoice         sql.NullString `db:"invoice"`
	CreatedAt       time.Time      `db:"created_at"`
	UpdatedAt       time.Time      `db:"updated_at"`
}
