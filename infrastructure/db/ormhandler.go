package db

import (
    "github.com/jinzhu/gorm"
    _  "github.com/go-sql-driver/mysql"
    "fmt"
	"migoV2/interfaces/repositories"
	"migoV2/bootstrap"
)

type OrmHandler struct {
    Conn *gorm.DB
}

func (handler *OrmHandler) Execute(statement string, params ...interface{}) (error) {
    return handler.Conn.Exec(statement, params...).Error
}

func (handler *OrmHandler) Query(tablename string, statement string, columns string, params ...interface{}) repositories.Row {
	rows, _ := handler.Conn.Table(tablename).Where(statement, params...).Select(columns).Rows()

    return rows
}

func NewOrmHandler() *OrmHandler {
	var adapter string
	var conn *gorm.DB

	adapter = bootstrap.App.DBConfig.String("adapter")

	switch adapter {
		case "mysql":
			conn = mysqlConn()
			break
		case "postgres":
			//postgresConn()
			break
	}

    ormHandler := new(OrmHandler)
    ormHandler.Conn = conn
    return ormHandler
}

func mysqlConn() (*gorm.DB){
	var (
		db *gorm.DB
		connectionString string
		err              error
	)
	
	connectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", bootstrap.App.DBConfig.String("username"), bootstrap.App.DBConfig.String("password"), bootstrap.App.DBConfig.String("host"), bootstrap.App.DBConfig.String("port"), bootstrap.App.DBConfig.String("database"))
	
	if db, err = gorm.Open("mysql", connectionString); err != nil {
		panic(err)
	}
	if err = db.DB().Ping(); err != nil {
		panic(err)
	}

	db.LogMode(true)
	db.DB().SetMaxIdleConns(bootstrap.App.DBConfig.Int("idle_conns"))
	db.DB().SetMaxOpenConns(bootstrap.App.DBConfig.Int("open_conns"))

	return db
}
