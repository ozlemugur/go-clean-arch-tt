package usecase

import (
	"context"

	"github.com/ozlemugur/go-clean-arch-tt/pkg/logger"
)

type AutoMessageUseCase struct {
	t Messager
	s AutoMessageSchedulerHandler
	l logger.Interface
}

// New -.
func NewAutoMessageUseCase(t Messager, s AutoMessageSchedulerHandler, l logger.Interface) *AutoMessageUseCase {
	return &AutoMessageUseCase{t, s, l}
}

func (uc *AutoMessageUseCase) StartAutoMessageSender(ctx context.Context) error {
	uc.s.Start()
	return nil
}

func (uc *AutoMessageUseCase) StopAutoMessageSender(ctx context.Context) error {
	uc.s.Stop()
	return nil
}
