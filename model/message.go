package model

import (
	"go-live-chat/config"
	"time"
)

//  msg struct represent a chat msg

type Message struct {
	ID        int       `json:"id"`
	SenderID  int       `json:"sender_id"`
	ReceiverID int      `json:"receiver_id"`
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
}

// save a new msg to db
func (m * Message) Save() error {
	_, err := config.DB.Exec(
		"INSERT INTO messages (sender_id, receiver_id, content, timestamp) VALUES (?, ?, ?, ?)",
		m.SenderID, m.ReceiverID, m.Content, m.Timestamp,
	)
	return err
}

//  getChat History retrieves chat
func GetChatHistory(userID1, userID2 int) ([]Message, error) {
	rows, err := config.DB.Query(
		`SELECT id, sender_id, receiver_id, content, timestamp 
		FROM messages 
		WHERE (sender_id = ? AND receiver_id = ?) 
			OR (sender_id = ? AND receiver_id = ?) 
		ORDER BY timestamp ASC`,
		userID1, userID2, userID2, userID1,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []Message
	for rows.Next() {
		var msg Message
		if err := rows.Scan(&msg.ID, &msg.SenderID, &msg.ReceiverID, &msg.Content, &msg.Timestamp); err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}
	return messages, nil
}
