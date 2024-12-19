package models

type Subscriber struct {
	ID   int    `db:"s_id"`
	Name string `db:"s_name"`
}