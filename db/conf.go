package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func CreateConn() (*sql.DB, error) {
	connectionUri := "root:opa123@tcp(127.0.0.1:3306)/test2"
	db, e := sql.Open("mysql", connectionUri)
	if e != nil{
		return nil, e
	}
	if e = db.Ping(); e != nil{
		return nil, e
	}
	return db, nil
}
