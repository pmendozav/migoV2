package domain

type BongoUserRepository interface {
	// FindUserByCredentials(userid, pass string) (BongoUser, error)
	CreateUser(userid, pass, firstname, lastname, email string) (error)
}

type BongoUser struct {
	Firstname  string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email  string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}