package dispatcher

import (
	"context"
	"fmt"

	"github.com/nick6969/go-clean-project/internal/domain"
	"github.com/nick6969/go-clean-project/internal/logger"
)

type Service struct {
	listeners map[domain.EventName][]domain.Listener
	logger    logger.Logger
}

func NewService(logger logger.Logger) *Service {
	return &Service{
		listeners: make(map[domain.EventName][]domain.Listener),
		logger:    logger,
	}
}

// dispatch 將事件傳遞給所有訂閱者
func (s *Service) dispatch(ctx context.Context, name domain.EventName, payload any) error {
	listeners, ok := s.listeners[name]
	if !ok {
		return fmt.Errorf("no listeners for event: %s", name)
	}

	for _, listener := range listeners {
		if err := listener.Handle(ctx, payload); err != nil {
			return err
		}
	}
	return nil
}

// RegisterListener 註冊事件監聽器
func (s *Service) RegisterListener(name domain.EventName, listener domain.Listener) {
	s.listeners[name] = append(s.listeners[name], listener)
}

// DispatchUserRegistered 發布 UserRegistered 事件
func (s *Service) DispatchUserRegistered(ctx context.Context, userID int, email string) {
	go func() {
		err := s.dispatch(ctx, domain.EventUserRegistered, domain.EventUserRegisteredPayload{
			UserID: userID,
			Email:  email,
		})
		if err != nil {
			s.logger.With(ctx).Error(ctx, "failed to dispatch user registered event", err)
		}
	}()
}
