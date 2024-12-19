package repository

import "github.com/jmoiron/sqlx"

type Generat interface {
	GenerateAuthors(db *sqlx.DB, count int)
	GenerateBooks(db *sqlx.DB, count int)
	GenerateGenres(db *sqlx.DB, count int)
	GenerateSubscribers(db *sqlx.DB, count int)
	GenerateBooksAuthors(db *sqlx.DB, count int)
	GenerateBooksGenres(db *sqlx.DB, count int)
	GenerateSubscriptions(db *sqlx.DB, count int)
}

type Repository struct {
	Generator Generat
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Generator: NewGenerator(db),
	}
}