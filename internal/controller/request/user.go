package request

type Register struct {
	Email    string `json:"email" binding:"required,email" example:"user@example.com"`
	Password string `json:"password" binding:"required,min=8,max=32" example:"P@ssw0rd!"`
}

type Login struct {
	Email    string `json:"email" binding:"required,email" example:"user@example.com"`
	Password string `json:"password" binding:"required,min=8,max=32" example:"P@ssw0rd!"`
}

type ChangePassword struct {
	Password    string `json:"password" binding:"required,min=8,max=32" example:"P@ssw0rd!"`
	NewPassword string `json:"new_password" binding:"required,min=8,max=32,nefield=Password" example:"N3wP@ssw0rd!"`
}
