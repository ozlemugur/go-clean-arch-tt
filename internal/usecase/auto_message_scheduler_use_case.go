package usecase

import (
	"context"

	"github.com/ozlemugur/go-clean-arch-tt/pkg/logger"
)

type AutoMessageSchedulerUseCase struct {
	t Messager
	s AutoMessageSchedulerHandler
	l logger.Interface
}

// New -.
func NewAutoMessageSchedulerUseCase(t Messager, s AutoMessageSchedulerHandler, l logger.Interface) *AutoMessageSchedulerUseCase {
	return &AutoMessageSchedulerUseCase{t, s, l}
}

func (uc *AutoMessageSchedulerUseCase) StartAutoMessageSender(ctx context.Context) error {
	uc.s.Start()
	return nil
}

func (uc *AutoMessageSchedulerUseCase) StopAutoMessageSender(ctx context.Context) error {
	uc.s.Stop()
	return nil
}
