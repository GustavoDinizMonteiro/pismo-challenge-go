package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"pismo-challenge-go/util"
)

func CreateAccount(w http.ResponseWriter, r *http.Request, conn *sql.DB) {
	body := util.ParseBody(r.Body)
	documentNumber := body["document_number"]

	_, error := conn.Exec("INSERT INTO tb_account (document_number) VALUES(?)", documentNumber)
	if error != nil {
		panic(error.Error())
	}

	fmt.Fprintf(w, "New post was created")
}
