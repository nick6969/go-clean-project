package token

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/pascaldekloe/jwt"
)

type Service struct {
	privateKey *ecdsa.PrivateKey
}

func NewService(key []byte) (*Service, error) {
	block, _ := pem.Decode(key)

	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block")
	}

	privatekey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse EC private key: %w", err)
	}

	return &Service{privateKey: privatekey}, nil
}

// signToken 使用私鑰對 Claims 進行簽名，並返回簽名後的 JWT 字串
func (s *Service) signToken(claims jwt.Claims) (string, error) {
	signed, err := claims.ECDSASign("ES512", s.privateKey)
	if err != nil {
		return "", err
	}

	return string(signed), nil
}

// checkToken 使用公鑰驗證 JWT 字串的簽名，並返回解析後的 Claims
func (s *Service) checkToken(token string) (*jwt.Claims, error) {
	return jwt.ECDSACheck([]byte(token), &s.privateKey.PublicKey)
}

// 使用固定的結構來生成不同類型的 Token
func (s *Service) generateToken(tokenType tokenType, userID int, expire time.Duration) (string, error) {
	var claims jwt.Claims
	random := uuid.New()

	now := time.Now().Round(time.Second)
	expires := now.Add(expire).Round(time.Second)

	claims.Issued = jwt.NewNumericTime(now)
	claims.Expires = jwt.NewNumericTime(expires)
	claims.Set = map[string]any{
		"uid":  random.String(),
		"id":   userID,
		"type": tokenType,
	}

	return s.signToken(claims)
}

// 驗證 Token 的有效性，並返回 Token 中的使用者 ID
func (s *Service) validateToken(tokenType tokenType, token string) (int, error) {
	claims, err := s.checkToken(token)
	if err != nil {
		return 0, fmt.Errorf("invalid token: %w", err)
	}

	if claims.Expires.Time().Before(time.Now()) {
		return 0, fmt.Errorf("token expired")
	}

	if claims.Set["type"] != string(tokenType) {
		return 0, fmt.Errorf("invalid token type")
	}

	userID, ok := claims.Set["id"].(float64)
	if !ok {
		return 0, fmt.Errorf("invalid user ID in token")
	}

	return int(userID), nil
}
