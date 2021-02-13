package handlers

import (
	"fmt"
	"log"
	"net/http"
	"pismo-challenge-go/db"
	"pismo-challenge-go/util"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	conn, error := db.CreateConn()
	if error != nil {
		log.Println("Database connection refused")
		w.Write([]byte("Error"))
	}
	defer conn.Close()

	stmt, error := conn.Prepare("INSERT INTO tb_account (id, document_number, credit_limit) VALUES(?, ?, ?)")
	if error != nil {
		panic(error.Error())
	}

	body := util.ParseBody(r.Body)
	documentNumber := body["document_number"]

	_, error = stmt.Exec(8, documentNumber, 0)
	if error != nil {
		panic(error.Error())
	}

	fmt.Fprintf(w, "New post was created")
}
