package repository

import (
	"base/pkg/db"
	"base/pkg/log"
	"base/service/messages/model"
)

// GetAllMessages : Get All Messages From Database
func GetAllMessages() ([]model.Message, error) {
	rows, err := db.PSQL.Query("SELECT id, content, created_at FROM messages")

	if err != nil {
		log.Println(log.LogLevelError, "query-get-all-messages", err.Error())
		return nil, err
	}
	defer rows.Close()

	var messages []model.Message
	for rows.Next() {
		var mess model.Message
		if err := rows.Scan(&mess.ID, &mess.Content, &mess.CreatedAt); err != nil {
			log.Println(log.LogLevelError, "scan-row-to-message", err.Error())
			return nil, err
		}
		messages = append(messages, mess)
	}

	return messages, nil
}
