package repository

import (
	"log"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"golang.org/x/exp/rand"
)

type Generator struct {
	db *sqlx.DB
}

func NewGenerator(db *sqlx.DB) *Generator {
	return &Generator{db: db}
}

func (r *Generator) GenerateAuthors(db *sqlx.DB, count int) {
	for i := 0; i < count; i++ {
		name := gofakeit.BookAuthor()
		_, err := db.Exec("INSERT INTO library.authors (a_name) VALUES ($1)", name)
		if err != nil {
			log.Printf("failed to insert author: %v", err)
		}
	}
		logrus.Infof("Generated %d authors", count)
}

func (r *Generator) GenerateBooks(db *sqlx.DB, count int) {
	for i := 0; i < count; i++ {
		name := gofakeit.BookTitle()
		year := rand.Intn(2023-1500) + 1500
		quantity := rand.Intn(100) + 1
		_, err := db.Exec("INSERT INTO library.books (b_name, b_year, b_quantity) VALUES ($1, $2, $3)", name, year, quantity)
		if err != nil {
			log.Printf("failed to insert book: %v", err)
		}
	}
	logrus.Infof("Generated %d books", count)
}

func (r *Generator) GenerateGenres(db *sqlx.DB, count int) {
	genres := []string{
		"Adventure", "Autobiography", "Biography", "Classic", "Comedy",
		"Coming-of-Age", "Crime", "Dark Fantasy", "Dystopian", "Epic",
		"Erotica", "Fantasy", "Fiction", "Gothic", "Graphic Novel",
		"Historical Fiction", "Horror", "Humor", "Inspirational", "Literary Fiction",
		"Magic Realism", "Memoir", "Military Fiction", "Mystery", "Mythology",
		"Non-fiction", "Paranormal", "Philosophy", "Poetry", "Political Fiction",
		"Psychological Fiction", "Romance", "Satire", "Science Fiction", "Self-help",
		"Short Stories", "Spirituality", "Sports Fiction", "Supernatural", "Suspense",
		"Thriller", "Time Travel", "Tragedy", "Travel", "Urban Fantasy",
		"War", "Western", "Women's Fiction", "Young Adult", "Children's Fiction",
		"Action", "Adventure Romance", "Alien Invasion", "Alternate History", "Animal Stories",
		"Anthology", "Apocalyptic", "Art Books", "Audiobook", "Chick Lit",
		"Christian Fiction", "Cozy Mystery", "Cyberpunk", "Dark Comedy", "Detective Fiction",
		"Drama", "Epic Fantasy", "Fairy Tale", "Family Saga", "Feminist Fiction",
		"Folklore", "Game Lit", "Gothic Romance", "Hard Science Fiction", "High Fantasy",
		"Historical Mystery", "Historical Romance", "Holocaust Fiction", "Inspirational Romance", "Legal Thriller",
		"Low Fantasy", "Martial Arts Fiction", "Media Tie-in", "Medical Fiction", "Modern Classic",
		"Occult", "Outdoors", "Parody", "Political Thriller", "Pop Culture",
		"Post-Apocalyptic", "Religious Fiction", "Reverse Harem", "Science Non-fiction", "Slice of Life",
		"Space Opera", "Steampunk", "Survival Fiction", "Technology", "Victorian Fiction",
		"Virtual Reality", "Zombie Fiction",
	}
	for i := 0; i < count; i++ {
		name := genres[i]
    _, err := db.Exec("INSERT INTO library.genres (g_name) VALUES ($1)", name)
    if err != nil {
        log.Printf("failed to insert genre: %v", err)
    }
}
logrus.Infof("Generated %d genres", count)
}

func (r *Generator) GenerateSubscribers(db *sqlx.DB, count int) {
	for i := 0; i < count; i++ {
		name := gofakeit.Name()
		_, err := db.Exec("INSERT INTO library.subscribers (s_name) VALUES ($1)", name)
		if err != nil {
			log.Printf("failed to insert subscriber: %v", err)
		}
	}
	logrus.Infof("Generated %d subscribers", count)
}

func (r *Generator) GenerateBooksAuthors(db *sqlx.DB, count int) {
	for i := 0; i < count; i++ {
		bookID := rand.Intn(100000) + 1
		authorID := rand.Intn(10000) + 1
		_, err := db.Exec("INSERT INTO library.m2m_books_authors (b_id, a_id) VALUES ($1, $2)", bookID, authorID)
		if err != nil {
			log.Printf("failed to insert book-author link: %v", err)
		}
	}
	logrus.Infof("Generated %d book-author links", count)
}

func (r *Generator) GenerateBooksGenres(db *sqlx.DB, count int) {
	for i := 0; i < count; i++ {
		bookID := rand.Intn(100000) + 1
		genreID := rand.Intn(100) + 1
		_, err := db.Exec("INSERT INTO library.m2m_books_genres (b_id, g_id) VALUES ($1, $2)", bookID, genreID)
		if err != nil {
			log.Printf("failed to insert book-genre link: %v", err)
		}
	}
	logrus.Infof("Generated %d book-genre links", count)
}

func (r *Generator) GenerateSubscriptions(db *sqlx.DB, count int) {
	for i := 0; i < count; i++ {
		subscriberID := rand.Intn(1000000) + 1
		bookID := rand.Intn(100000) + 1
		start := gofakeit.DateRange(time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC), time.Now())
		finish := gofakeit.DateRange(start, time.Now().AddDate(1, 0, 0))
		status := "Y"
		if rand.Intn(2) == 0 {
			status = "N"
		}
		_, err := db.Exec("INSERT INTO library.subscriptions (sb_subscriber, sb_book, sb_start, sb_finish, sb_is_active) VALUES ($1, $2, $3, $4, $5)", subscriberID, bookID, start, finish, status)
		if err != nil {
			log.Printf("failed to insert subscription: %v", err)
		}
	}
	logrus.Infof("Generated %d subscriptions", count)
}