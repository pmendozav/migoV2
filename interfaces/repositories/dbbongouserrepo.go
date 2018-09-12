package repositories

import (
	_ "mio2/domain"
	_ "errors"
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

// func (repo *DbBongoUserRepo) FindUserByCredentials(userid, pass string) (domain.BongoUser, error) {
// 	tablename := "usuarios"
// 	statement := "userid = ? and pass = ?"
// 	columns := "firstname, pass2"
// 	param1 := userid
// 	param2 := pass
// 	row := repo.dbHandler.Query(tablename, statement, columns, param1, param2)

// 	var (
// 		firstname string
// 		pass2 string
// 	)

// 	if ok := row.Next(); !ok {
// 		return domain.BongoUser{}, errors.New("Invalid credentials!")
// 	}

//     row.Scan(&firstname, &pass2)
//     return domain.BongoUser{Firstname: firstname, Password: pass2}, nil
// }
