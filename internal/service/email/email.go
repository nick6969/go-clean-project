package email

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) SendEmail(email string, body any) error {
	// TODO: - Implement email sending logic here
	return nil
}
