package domain

const (
	EventUserRegistered EventName = "UserRegistered"
)

type EventUserRegisteredPayload struct {
	UserID int
	Email  string
}
