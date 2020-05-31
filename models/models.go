package models

import "time"

type Message struct {
	ID        string
	ChannelID string
	UserID    string
	Body      string
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
