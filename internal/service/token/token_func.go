package token

import (
	"time"
)

func (s *Service) GenerateAccessToken(userID int) (string, error) {
	return s.generateToken(tokenTypeAccessToken, userID, time.Hour*24)
}

func (s *Service) ValidateAccessToken(token string) (int, error) {
	return s.validateToken(tokenTypeAccessToken, token)
}
