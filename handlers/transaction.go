package handlers

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"pismo-challenge-go/db"
	"pismo-challenge-go/models"
	"pismo-challenge-go/util"
	"time"
)

type Account struct {
	id int `json:"id"`
	creditLimit float64 `json:"credit_limit"`
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	transaction := util.ParseTransaction(r.Body)
	accountId := transaction.AccountId
	account, err := getAccount(accountId)

	if err != nil {
		log.Println("Database connection refused", err)
		w.Write([]byte("Error"))
	}

	amount := transaction.Amount
	if transaction.OperationTypeId == 4 && account.creditLimit < math.Abs(amount) {
		w.Write([]byte("No credit available"))
	}

	account.creditLimit += amount
	saveTransaction(transaction, accountId)
	updateAccount(accountId, account.creditLimit)
	fmt.Fprintf(w, "New post was created")
}

func getAccount(accountId int) (Account, error) {
	conn, err := db.CreateConn()
	if err != nil {
		return Account{}, err
	}
	result, err := conn.Query("SELECT id, credit_limit FROM tb_account where id=?", accountId)
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

func updateAccount(accountId int, amount float64) {
	conn, error := db.CreateConn()
	if error != nil {
		panic(error.Error())
	}
	defer conn.Close()

	stmt, error := conn.Prepare("UPDATE tb_account SET credit_limit=? WHERE id=?")
	if error != nil {
		panic(error.Error())
	}

	_, error = stmt.Exec(amount, accountId)
	if error != nil {
		panic(error.Error())
	}
}

func saveTransaction(transaction models.Transaction, accountId int) {
	conn, error := db.CreateConn()
	if error != nil {
		panic(error.Error())
	}
	defer conn.Close()

	stmt, error := conn.Prepare("INSERT INTO tb_transaction (id, amount, account_id, operation_type_id, event_date) VALUES(?, ?, ?, ?, ?)")
	if error != nil {
		panic(error.Error())
	}

	amount := transaction.Amount
	operationId := transaction.OperationTypeId
	eventDate := time.Now()

	_, error = stmt.Exec(33, amount, accountId, operationId, eventDate)
	if error != nil {
		panic(error.Error())
	}
}