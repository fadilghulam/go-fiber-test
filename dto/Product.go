package dto

type Product struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Code  string `db:"code"`
	Order int    `db:"order"`
}
