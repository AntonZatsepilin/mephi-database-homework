package models

type Book struct {
	ID     int    `db:"b_id"`
	Name   string `db:"b_name"`
	Year   int    `db:"b_year"`
	Quantity int `db:"b_quantity"`
}