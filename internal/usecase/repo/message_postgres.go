package repo

import (
	"context"
	"fmt"

	"github.com/ozlemugur/go-clean-arch-tt/internal/entity"
	"github.com/ozlemugur/go-clean-arch-tt/pkg/postgres"
)

type MessageRepo struct {
	*postgres.Postgres
}

// New -.
func NewMessage(pg *postgres.Postgres) *MessageRepo {
	return &MessageRepo{pg}
}

func (r *MessageRepo) GetSentMessages(ctx context.Context) ([]entity.Message, error) {
	// SQL query to fetch messages with status 'sent'
	sql, _, err := r.Builder.
		Select("id, content, recipient_phone, status, created_at").
		From("messages").
		Where("status = ?", "sent").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("MessageRepo - GetSentMessages - r.Builder: %w", err)
	}
	rows, err := r.Pool.Query(ctx, sql, "sent")
	if err != nil {
		return nil, fmt.Errorf("MessageRepo - GetSentMessages - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	// Initialize the slice to store results
	messages := make([]entity.Message, 0)

	// Iterate over the rows and scan the results
	for rows.Next() {
		msg := entity.Message{}

		err = rows.Scan(&msg.ID, &msg.Content, &msg.RecipientPhone, &msg.Status, &msg.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("MessageRepo - GetSentMessages - rows.Scan: %w", err)
		}

		messages = append(messages, msg)
	}

	return messages, nil
}

func (r *MessageRepo) GetTwoUnSentMessages(ctx context.Context) ([]entity.Message, error) {
	// SQL query to fetch messages with status 'sent'
	sql, _, err := r.Builder.
		Select("id, content, recipient_phone, status, created_at").
		From("messages").
		Where("status = ?", "unsent").
		Limit(2).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("MessageRepo - GetUnSentMessages - r.Builder: %w", err)
	}

	// Execute the query
	rows, err := r.Pool.Query(ctx, sql, "unsent")
	if err != nil {
		return nil, fmt.Errorf("MessageRepo - GetUnSentMessages - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	// Initialize the slice to store results
	messages := make([]entity.Message, 0)

	// Iterate over the rows and scan the results
	for rows.Next() {
		msg := entity.Message{}

		err = rows.Scan(&msg.ID, &msg.Content, &msg.RecipientPhone, &msg.Status, &msg.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("MessageRepo - GetSentMessages - rows.Scan: %w", err)
		}

		messages = append(messages, msg)
	}

	return messages, nil
}

func (r *MessageRepo) InsertMessage(ctx context.Context, msg entity.Message) error {
	sql, args, err := r.Builder.
		Insert("messages").
		Columns("content", "recipient_phone").
		Values(msg.Content, msg.RecipientPhone).
		ToSql()
	if err != nil {
		return fmt.Errorf("MessageRepo - InsertMessage - r.Builder: %w", err)
	}
	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("MessageRepo - InsertMessage - r.Pool.Exec: %w", err)
	}

	return nil
}

func (r *MessageRepo) UpdateMessageStatus(ctx context.Context, messageID int, status string) error {
	sql, args, err := r.Builder.
		Update("messages").
		Set("status", status).
		Where("id = ?", messageID).
		ToSql()
	if err != nil {
		return fmt.Errorf("MessageRepo - UpdateMessageStatus - r.Builder: %w", err)
	}
	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("MessageRepo - UpdateMessageStatus - r.Pool.Exec: %w", err)
	}
	return nil
}
