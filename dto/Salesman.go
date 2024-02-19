package dto

import (
	"time"
)

type Salesman struct {
	ID        int       `db:"id"`
	BranchID  int       `db:"branch_id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
