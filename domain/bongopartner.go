package domain

type BongoCommonRepository interface {
	FindPartner(wildcard string) ([]BongoPartner, error)
}

type BongoPartner struct {
	IdPartner  string `json:"idpartner"`
	Name string `json:"name"`
	Key string `json:"key"`
	Code string `json:"code"`
}