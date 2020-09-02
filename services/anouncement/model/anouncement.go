package model

import (
	"time"
)

type Anouncement struct {
	Title     string    `json:"title" bson:"title"`
	Message   string    `json:"message" bson:"message"`
	Author    string    `json:"author" bson:"author"`
	Type      string    `json:"type" bson:"type"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
