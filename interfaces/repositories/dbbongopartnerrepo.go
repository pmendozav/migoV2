package repositories

import (
	"bytes"
	"migoV2/domain"
	_ "errors"
	_ "fmt"
)

type DbBongoCommonRepo DbRepo

func NewDbBongoCommonRepo(dbHandlers map[string]DbHandler) *DbBongoCommonRepo {
    DbBongoCommonRepo := new(DbBongoCommonRepo)
    DbBongoCommonRepo.dbHandlers = dbHandlers
    DbBongoCommonRepo.dbHandler = dbHandlers["DbBongoCommonRepo"]
    return DbBongoCommonRepo
}

func (repo *DbBongoCommonRepo) FindPartner(wildcard string) ([]domain.BongoPartner, error) {

	var b bytes.Buffer

	b.WriteString("%")
	b.WriteString(wildcard)
	b.WriteString("%")

	tablename := "partners"
	statement := "name like ?"
	columns := "idpartner, name, code, `key`"
	param1 := b.String()

	rows := repo.dbHandler.Query(tablename, statement, columns, param1)

	var (
		_idpartner string
		_name string
		_code string
		_key string
	)

	partners := make([]domain.BongoPartner, 0)

	for rows.Next() {
		rows.Scan(&_idpartner, &_name, &_code, &_key)
		partners = append(partners, domain.BongoPartner{IdPartner:_idpartner, Name:_name, Code:_code, Key:_key})
	}

	return partners, nil
}