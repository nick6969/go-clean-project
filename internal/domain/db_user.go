package domain

type DBUserModel struct {
	id           int
	email        string
	passwordHash string
}

func NewDBUserModel(id int, email, passwordHash string) *DBUserModel {
	return &DBUserModel{
		id:           id,
		email:        email,
		passwordHash: passwordHash,
	}
}

func (u *DBUserModel) ID() int {
	return u.id
}

func (u *DBUserModel) Email() string {
	return u.email
}

func (u *DBUserModel) PasswordHash() string {
	return u.passwordHash
}
