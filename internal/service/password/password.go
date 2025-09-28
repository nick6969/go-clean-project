package password

import (
	"errors"

	"github.com/nick6969/go-clean-project/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

// 提供密碼雜湊功能
func (s *Service) Hash(password string) (string, *domain.GPError) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", domain.NewGPErrorWithError(domain.ErrCodeInternalServer, err).Append("failed to hash password")
	}
	return string(hashedBytes), nil
}

// 提供密碼比對功能
func (s *Service) Compare(hashed, password string) *domain.GPError {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return domain.NewGPError(domain.ErrCodeInvalidPassword)
		}
		return domain.NewGPErrorWithError(domain.ErrCodeInternalServer, err).Append("failed to compare password")
	}
	return nil
}
