package login

import (
    "errors"
	_ "fmt"
)

func (interactor *BongoUserInteractor) ValidateUser(user *User) (error) {
	if !userFieldsValidation(*user) {
		return errors.New("Not valid fields")
	}

	tmp, err := interactor.BongoUserRepository.GetUserInfo(user.Username, user.Password)


	
	if err == nil {
		user.Firstname = tmp.Firstname
		user.Lastname = tmp.Lastname
		user.Email = tmp.Email
		user.Username = tmp.Username
	} else {
		user = nil
		return err
	}

	return nil
}
