package common

import (
    "migoV2/domain"
    _ "fmt"
)

type Logger interface {
    Log(message string) error
}

type Partner struct {
	IdPartner  string `json:"idpartner"`
	Name string `json:"name"`
	Key string `json:"key"`
	Code string `json:"code"`
}

type BongoCommonInteractor struct {
    BongoCommonRepository domain.BongoCommonRepository
    Logger Logger
}


// func userFieldsValidation(user User) (bool) {
// 	if len(user.Username) == 0 {
// 		return false
// 	}

// 	if len(user.Password) < 4 {
// 		return false
// 	}

// 	return true
// }