package models

import "time"

type Subscription struct {
	ID   int    `db:"sb_id"`
	SubscriberID int `db:"sb_subscriber"`
	SubscriberBookID int `db:"sb_book"`
	SubscriberStart time.Time `db:"sb_start"`
	SubscriberFinish time.Time `db:"sb_finish"`
	SubscriberStatus string `db:"sb_is_active"`
}
