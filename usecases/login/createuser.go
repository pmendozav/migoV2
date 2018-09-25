package login

import (
    "errors"
    _ "fmt"
)

func (interactor *BongoUserInteractor) CreateUser(user User) (error) {
	if !userFieldsValidation(user) {
		return errors.New("Not valid fields")
	}
	return interactor.BongoUserRepository.CreateUser(user.Username, user.Password, user.Firstname, user.Lastname, user.Email)
}
