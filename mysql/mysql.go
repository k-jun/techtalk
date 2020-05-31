package mysql

import (
	"database/sql"
	"techtalk/models"

	_ "github.com/go-sql-driver/mysql"
)

type IMySQL interface {
	GetChannelMessage(cid string) ([]models.Message, error)
	CreateChannelMessage(cid string, m *models.Message) error
	UpdateChannelMessage(cid string, m *models.Message) error
	DeleteChannelMessage(cid string, mid string) error
}

type sMySQL struct {
	db *sql.DB
}

func NewSMySQL(conn *sql.DB) IMySQL {
	return &sMySQL{
		db: conn,
	}
}

func (m *sMySQL) GetChannelMessage(channelID string) ([]models.Message, error) {
	messages := make([]models.Message, 0)
	return messages, nil
}

func (m *sMySQL) CreateChannelMessage(channelID string, message *models.Message) error {
	return nil
}

func (m *sMySQL) UpdateChannelMessage(channelID string, message *models.Message) error {
	return nil
}

func (m *sMySQL) DeleteChannelMessage(channelID string, messageID string) error {
	return nil
}
