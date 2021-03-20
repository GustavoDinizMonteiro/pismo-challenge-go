package handlers

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math"
	"net/http"
	"pismo-challenge-go/models"
	"pismo-challenge-go/util"
	"time"
)

type Account struct {
	id int `json:"id"`
	creditLimit float64 `json:"credit_limit"`
}

func CreateTransaction(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	transaction := util.ParseTransaction(r.Body)
	accountId := transaction.AccountId

	conn, _ := db.Conn(context.Background())
	defer conn.Close()

	account, err := getAccount(accountId, conn)
	if err != nil {
		log.Println("Database connection refused", err)
		w.Write([]byte("Error"))
	}

	amount := transaction.Amount
	if transaction.OperationTypeId == 4 && account.creditLimit < math.Abs(amount) {
		w.Write([]byte("No credit available"))
	}

	account.creditLimit += amount
	saveTransaction(transaction, accountId, conn)
	updateAccount(accountId, account.creditLimit, conn)
	fmt.Fprintf(w, "New post was created")
}

func getAccount(accountId int, conn *sql.Conn) (Account, error) {
	result, err := conn.QueryContext(context.Background(),"SELECT id, credit_limit FROM tb_account where id=?", accountId)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var account Account
	for result.Next() {
		err := result.Scan(&account.id, &account.creditLimit)
		if err != nil {
			panic(err.Error())
		}
	}
	return account, nil
}

func updateAccount(accountId int, amount float64, conn *sql.Conn) {
	stmt, error := conn.PrepareContext(context.Background(),"UPDATE tb_account SET credit_limit=? WHERE id=?")
	if error != nil {
		panic(error.Error())
	}

	_, error = stmt.Exec(amount, accountId)
	if error != nil {
		panic(error.Error())
	}
	stmt.Close()
}

func saveTransaction(transaction models.Transaction, accountId int, conn *sql.Conn) {
	stmt, error := conn.PrepareContext(context.Background(),"INSERT INTO tb_transaction (amount, account_id, operation_type_id, event_date) VALUES(?, ?, ?, ?)")
	if error != nil {
		panic(error.Error())
	}

	amount := transaction.Amount
	operationId := transaction.OperationTypeId
	eventDate := time.Now()

	_, error = stmt.Exec(amount, accountId, operationId, eventDate)
	if error != nil {
		panic(error.Error())
	}
	stmt.Close()
}