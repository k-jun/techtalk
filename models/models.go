package models

type Message struct {
	ID        string
	ChannelID string
	UserID    string
	Body      string
	Type      string
	CreatedAt int64
	UpdatedAt int64
}
