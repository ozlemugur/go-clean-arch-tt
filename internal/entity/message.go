package entity

import "time"

// Message -.
type Message struct {
	ID             int       `json:"-" db:"id"`
	Content        string    `json:"content" db:"content"`
	RecipientPhone string    `json:"recipient_phone" db:"recipient_phone"`
	Status         string    `json:"-" db:"status"`
	CreatedAt      time.Time `json:"-" db:"created_at"`
}

type AutoMessageSender struct {
	Action string `json:"action" example:"start/stop" `
}
