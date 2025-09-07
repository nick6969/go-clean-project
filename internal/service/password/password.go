package password

import "golang.org/x/crypto/bcrypt"

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

// 提供密碼雜湊功能
func (s *Service) Hash(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

// 提供密碼比對功能
func (s *Service) Compare(hashed, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
}
