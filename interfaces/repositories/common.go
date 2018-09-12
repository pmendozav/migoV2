package repositories

type DbHandler interface {
    Execute(statement string, params ...interface{}) (error)
    Query(tablename string, statement string, columns string, params ...interface{}) (Row)
}

type Row interface {
    Scan(params ...interface{}) (error)
    Next() (bool)
    Close() (error)
}

type DbRepo struct {
    dbHandlers map[string]DbHandler
    dbHandler  DbHandler
}