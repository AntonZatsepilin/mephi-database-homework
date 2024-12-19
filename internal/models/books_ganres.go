package models

type BooksGanres struct {
	BookID  int `db:"b_id"`
	GenreID int `db:"g_id"`
}