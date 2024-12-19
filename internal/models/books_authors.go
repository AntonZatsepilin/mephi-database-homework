package models

type BooksAuthors struct {
	BookID   int `db:"b_id"`
	AuthorID int `db:"a_id"`
}