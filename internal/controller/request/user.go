package request

type Register struct {
	Email    string `json:"email" binding:"required,email" example:"user@example.com"`
	Password string `json:"password" binding:"required,min=8,max=32" example:"P@ssw0rd!"`
}

type Login struct {
	Email    string `json:"email" binding:"required,email" example:"user@example.com"`
	Password string `json:"password" binding:"required,min=8,max=32" example:"P@ssw0rd!"`
}
