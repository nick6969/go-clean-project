package token

import (
	"time"

	"github.com/nick6969/go-clean-project/internal/domain"
)

func (s *Service) GenerateAccessToken(userID int) (string, *domain.GPError) {
	token, err := s.generateToken(tokenTypeAccessToken, userID, time.Hour*24)
	if err != nil {
		return "", domain.NewGPErrorWithError(domain.ErrCodeInternalServer, err).Append("failed to generate access token")
	}
	return token, nil
}

func (s *Service) ValidateAccessToken(token string) (int, *domain.GPError) {
	userID, err := s.validateToken(tokenTypeAccessToken, token)
	if err != nil {
		return 0, domain.NewGPErrorWithError(domain.ErrCodeInternalServer, err).Append("failed to validate access token")
	}
	return userID, nil
}
