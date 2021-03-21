package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"pismo-challenge-go/handlers"
	"pismo-challenge-go/server"
	"time"
)

var db *sql.DB

func main() {
	server := server.CreateServer()

	connectionUri := "root:opa123@tcp(127.0.0.1:3306)/test2"
	db, e := sql.Open("mysql", connectionUri)
	if e != nil {
		panic(e.Error())
	}
	if e = db.Ping(); e != nil {
		panic(e.Error())
	}

	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Second)

	server.HandleFunc("/accounts", func(writer http.ResponseWriter, request *http.Request) {
		handlers.CreateAccount(writer, request, db)
	}).Methods("GET")

	server.HandleFunc("/transactions", func(writer http.ResponseWriter, request *http.Request) {
		handlers.CreateTransaction(writer, request, db)
	}).Methods("POST")
	http.ListenAndServe(":8080", server)
}
