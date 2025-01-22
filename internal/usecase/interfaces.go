// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"github.com/ozlemugur/go-clean-arch-tt/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (

	/*Message UseCase*/
	Messager interface {
		GetSentMessages(context.Context) ([]entity.Message, error)
		InsertMessage(context.Context, entity.Message) error
		GetTwoUnSentMessages(context.Context) ([]entity.Message, error)
		FetchAndSendMessages(context.Context) error
	}

	// MessageRepo - Message Handler
	MessageRepositoryHandler interface {
		GetSentMessages(context.Context) ([]entity.Message, error)
		GetTwoUnSentMessages(context.Context) ([]entity.Message, error)
		InsertMessage(context.Context, entity.Message) error
		UpdateMessageStatus(context.Context, int, string) error
	}

	/******/
	/*AutoMessage UseCase*/
	AutoMessager interface {
		StartAutoMessageSender(context.Context) error
		StopAutoMessageSender(context.Context) error
	}

	AutoMessageSchedulerHandler interface {
		Start()
		Stop()
		IsRunning() bool
	}

	// MessageSenderWebAPI -.
	MessageSenderWebAPIHandler interface {
		SendMessage(to, content string) (string, error)
	}
)
