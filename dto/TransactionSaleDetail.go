package dto

import (
	"time"
)

type TransactionSaleDetail struct {
	ID            int       `db:"id"`
	TransactionID int       `db:"transaction_id"`
	ProductID     int       `db:"product_id"`
	Price         int       `db:"price"`
	Qty           int       `db:"qty"`
	Discount      int       `db:"discount"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}
