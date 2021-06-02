package models

import "time"

/* InsertTweet is the structure to save Tweet in DB */
type InsertTweet struct {
	UserID    string    `bson:"userId" json:"userId,omitempty"`
	Message   string    `bson:"message" json:"message,omitempty"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt,omitempty"`
}
