package usecase

import (
	"context"
	"fmt"

	"github.com/ozlemugur/go-clean-arch-tt/internal/entity"
	"github.com/ozlemugur/go-clean-arch-tt/pkg/logger"
)

// MessageUseCase -.
type MessageUseCase struct {
	repo   MessageRepositoryHandler
	webApi MessageSenderWebAPIHandler
	l      logger.Interface
}

// New -.
func NewMessageUseCase(r MessageRepositoryHandler, w MessageSenderWebAPIHandler, l logger.Interface) *MessageUseCase {
	return &MessageUseCase{
		repo:   r,
		webApi: w,
		l:      l,
	}
}

// Retrieve all sent messages from database
func (uc *MessageUseCase) GetSentMessages(ctx context.Context) ([]entity.Message, error) {
	messages, err := uc.repo.GetSentMessages(ctx)
	if err != nil {
		return nil, fmt.Errorf("MessageUseCase - GetSentMessages - s.repo.GetMessages: %w", err)
	}
	return messages, nil
}

// Retrieve two unsent messages from database
func (uc *MessageUseCase) GetTwoUnSentMessages(ctx context.Context) ([]entity.Message, error) {
	messages, err := uc.repo.GetTwoUnSentMessages(ctx)
	if err != nil {
		return nil, fmt.Errorf("MessageUseCase - GetUnSentMessages - s.repo.GetMessages: %w", err)
	}
	return messages, nil
}

func (uc *MessageUseCase) InsertMessage(ctx context.Context, msg entity.Message) error {
	// Validate the message content
	if len(msg.Content) > 160 {
		return fmt.Errorf("MessageUseCase - InsertMessage: content exceeds character limit of 160")
	}
	// Insert the message using the repository
	err := uc.repo.InsertMessage(ctx, msg)
	if err != nil {
		return fmt.Errorf("MessageUseCase - InsertMessage - uc.repo.InsertMessage: %w", err)
	}
	return nil
}

// FetchAndSendMessages -.
func (uc *MessageUseCase) FetchAndSendMessages(ctx context.Context) error {
	/// we should fetch two messages from the db and sent to the sender
	// if we have problem while sending we should reconsider the solution of the problem to not lose the data.
	// if everything is okay, we should update the records.
	// we need transaction managemnet here.

	messages, err := uc.GetTwoUnSentMessages(ctx)
	if err != nil {
		return fmt.Errorf("MessageUseCase - ProcessUnsentMessages - GetTwoUnSentMessages: %w", err)
	}
	if len(messages) == 0 {
		uc.l.Info("No unsent messages left to process.")
		return nil
	}
	for _, msg := range messages {
		_, err := uc.webApi.SendMessage(msg.RecipientPhone, msg.Content)
		if err != nil {
			uc.l.Error(err, "Failed to send message", "MessageID", msg.ID)
			continue
		}

		err = uc.repo.UpdateMessageStatus(ctx, msg.ID, "sent")
		if err != nil {
			uc.l.Error(err, "Failed to update message status", "MessageID", msg.ID)
		}
	}
	return nil
}
