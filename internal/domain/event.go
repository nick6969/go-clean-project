package domain

import "context"

type EventName string

type Listener interface {
	Handle(ctx context.Context, payload any) error
}
