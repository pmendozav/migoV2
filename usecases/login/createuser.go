package login

import (
    "migoV2/domain"
    "errors"
    "fmt"
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

type BongoCreateUserInteractor struct {
    BongoUserRepository domain.BongoUserRepository
    Logger Logger
}

func (interactor *BongoCreateUserInteractor) CreateUser(user User) (error) {
	if !userFieldsValidation(user) {
		return errors.New("Not valid fields")
	}
	fmt.Println("-.-")
	return interactor.BongoUserRepository.CreateUser(user.Username, user.Password, user.Firstname, user.Lastname, user.Email)
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