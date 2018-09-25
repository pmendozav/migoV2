package repositories

import (
	"migoV2/domain"
	"errors"
	_ "fmt"
)

type DbBongoUserRepo DbRepo

func NewDbBongoUserRepo(dbHandlers map[string]DbHandler) *DbBongoUserRepo {
    dbBongoUserRepo := new(DbBongoUserRepo)
    dbBongoUserRepo.dbHandlers = dbHandlers
    dbBongoUserRepo.dbHandler = dbHandlers["DbBongoUserRepo"]
    return dbBongoUserRepo
}

func (repo *DbBongoUserRepo) CreateUser(userid, pass, firstname, lastname, email string) (error) {
	statement := "INSERT INTO usuarios (userid, pass, firstname, lastname, email) VALUES (?, ?, ?, ?, ?)"
	
	return repo.dbHandler.Execute(statement, userid, pass, firstname, lastname, email)
}

func (repo *DbBongoUserRepo) GetUserInfo(userid, pass string) (domain.BongoUser, error) {
	tablename := "usuarios"
	statement := "userid = ? and pass = ?"
	columns := "firstname, lastname, userid, email"
	param1 := userid
	param2 := pass

	rows := repo.dbHandler.Query(tablename, statement, columns, param1, param2)

	var (
		_firstname string
		_lastname string
		_userid string
		_email string
	)

	if ok := rows.Next(); !ok {
		return domain.BongoUser{}, errors.New("Invalid credentials!")
	}

	rows.Scan(&_firstname, &_lastname, &_userid, &_email)
	
    return domain.BongoUser{Firstname: _firstname, Lastname: _lastname, Username:_userid, Email:_email}, nil
}
