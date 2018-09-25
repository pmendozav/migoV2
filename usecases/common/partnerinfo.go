package common

import (
    _ "errors"
    _ "fmt"
)

func (interactor *BongoCommonInteractor) FindPartner(wildcard string) ([]Partner, error) {
	tmp, _ := interactor.BongoCommonRepository.FindPartner(wildcard)

	partners := make([]Partner, len(tmp))

	for i, p := range tmp {
		partners[i].IdPartner = p.IdPartner
		partners[i].Name = p.Name
		partners[i].Key = p.Key
		partners[i].Code = p.Code
    }

	return partners, nil
}
