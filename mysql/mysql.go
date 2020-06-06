package mysql

import (
	"database/sql"
	"errors"
	"strconv"
	"techtalk/models"

	_ "github.com/go-sql-driver/mysql"
)

type IMySQL interface {
	GetChannelMessage(cid string) ([]models.Message, error)
	CreateChannelMessage(m *models.Message) error
	UpdateChannelMessage(m *models.Message) error
	DeleteChannelMessage(mid string) error
}

type sMySQL struct {
	db *sql.DB
}

func NewSMySQL(conn *sql.DB) (IMySQL, error) {
	if err := conn.Ping(); err != nil {
		return nil, err
	}

	return &sMySQL{db: conn}, nil
}

func (m *sMySQL) GetChannelMessage(channelID string) ([]models.Message, error) {
	messages := make([]models.Message, 0)
	rows, err := m.db.Query(`
		SELECT id, channel_id, user_id, type, body, created_at, updated_at
		FROM messages as m
		WHERE m.channel_id = ?
		ORDER BY id DESC
		LIMIT 10`,
		channelID,
	)
	if err != nil {
		return messages, err
	}

	for rows.Next() {
		me := models.Message{}
		err = rows.Scan(&me.ID, &me.ChannelID, &me.UserID, &me.Type, &me.Body, &me.CreatedAt, &me.UpdatedAt)
		if err != nil {
			return messages, err
		}
		messages = append(messages, me)
	}

	return messages, nil
}

func (m *sMySQL) CreateChannelMessage(message *models.Message) error {

	result, err := m.db.Exec(`
		INSERT INTO messages(id, channel_id, user_id, type, body, created_at, updated_at)
		VALUES(NULL, ?, ?, ?, ?, ?, ?);`,
		message.ChannelID, message.UserID, message.Type, message.Body, message.CreatedAt, message.UpdatedAt,
	)
	if err != nil {
		return err
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	message.ID = strconv.FormatInt(lastID, 10)
	return nil
}

func (m *sMySQL) UpdateChannelMessage(message *models.Message) error {
	result, err := m.db.Exec(`
		UPDATE messages as m 
		SET type = ?,
			body = ?,
			updated_at = ?,
			created_at = ?
		WHERE m.id = ?;`,
		message.Type, message.Body, message.CreatedAt, message.UpdatedAt, message.ID,
	)
	if err != nil {
		return err
	}
	num, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if num != 1 {
		return errors.New("message not found or no-update for the record")
	}
	return nil
}

func (m *sMySQL) DeleteChannelMessage(messageID string) error {
	result, err := m.db.Exec(`
		DELETE FROM messages
		WHERE messages.id = ?;`,
		messageID,
	)
	if err != nil {
		return err
	}
	num, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if num != 1 {
		return errors.New("message not found")
	}
	return nil
}
