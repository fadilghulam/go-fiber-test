package dto

type Customer struct {
	ID           int    `db:"id"`
	SalesmanID   int    `db:"salesman_id"`
	CustomerName string `db:"customer_name"`
	OutletName   string `db:"outlet_name"`
	Type         string `db:"type"`
	Address      string `db:"address"`
	SubDistrict  string `db:"sub_district"`
	District     string `db:"district"`
	Province     string `db:"province"`
}
