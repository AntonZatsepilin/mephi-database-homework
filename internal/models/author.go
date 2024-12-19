package models

type Author struct {
	ID   int    `db:"a_id"`
	Name string `db:"a_name"`
}