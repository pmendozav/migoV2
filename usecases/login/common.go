package login

import (
    "migoV2/domain"
    _ "fmt"
)

type Logger interface {
    Log(message string) error
}

type User struct {
	Firstname  string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email  string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type BongoUserInteractor struct {
    BongoUserRepository domain.BongoUserRepository
    Logger Logger
}


func userFieldsValidation(user User) (bool) {
	if len(user.Username) == 0 {
		return false
	}

	if len(user.Password) < 4 {
		return false
	}

	return true
}