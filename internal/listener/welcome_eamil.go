package listener

import (
	"context"
	"fmt"

	"github.com/nick6969/go-clean-project/internal/domain"
)

type EmailService interface {
	SendEmail(email string, body any) error
}

type WelcomeEmail struct {
	emailService EmailService
}

func NewWelcomeEmail(emailService EmailService) *WelcomeEmail {
	return &WelcomeEmail{emailService: emailService}
}

func (w *WelcomeEmail) Handle(ctx context.Context, event any) error {
	userRegisteredEvent, ok := event.(domain.EventUserRegisteredPayload)
	if !ok {
		return fmt.Errorf("invalid event type")
	}

	// Send welcome email
	return w.emailService.SendEmail(userRegisteredEvent.Email, "Welcome!")
}
