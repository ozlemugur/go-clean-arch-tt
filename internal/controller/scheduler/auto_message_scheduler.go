package scheduler

import (
	"context"
	"sync"
	"time"

	"github.com/ozlemugur/go-clean-arch-tt/internal/usecase"
	"github.com/ozlemugur/go-clean-arch-tt/pkg/logger"
)

type AutoMessageScheduler struct {
	m       usecase.Messager
	l       logger.Interface
	cancel  context.CancelFunc
	running bool
	mu      sync.Mutex
}

func NewAutoMessageScheduler(m usecase.Messager, l logger.Interface) *AutoMessageScheduler {
	return &AutoMessageScheduler{
		m: m,
		l: l,
	}
}

// Start begins the automatic message sending process
func (s *AutoMessageScheduler) Start() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.running {
		s.l.Info("Scheduler is already running")
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	s.cancel = cancel
	s.running = true

	ticker := time.NewTicker(2 * time.Minute)

	// Start the scheduler goroutine
	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				err := s.m.FetchAndSendMessages(ctx)
				if err != nil {
					s.l.Error(err, "AutoMessageScheduler - FetchAndSendMessages failed")
					continue
				}
				s.l.Info("AutoMessageScheduler - Messages successfully fetched and sent")
			case <-ctx.Done():
				s.l.Info("AutoMessageScheduler - Stopping scheduler")
				return
			}
		}
	}()

	s.l.Info("AutoMessageScheduler - Started")
}

// Stop halts the automatic message sending process
func (s *AutoMessageScheduler) Stop() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.running {
		s.l.Info("Scheduler is not running")
		return
	}

	s.cancel()
	s.running = false

	s.l.Info("AutoMessageScheduler - Stopped")
}

// IsRunning checks if the scheduler is running
func (s *AutoMessageScheduler) IsRunning() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.running
}
