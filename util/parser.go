package util

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
)

func ParseBody(bodyReader io.ReadCloser) map[string]string {
	body, err := ioutil.ReadAll(bodyReader)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)

	log.Println(keyVal)
	return keyVal
}

type Transaction struct {
	AccountId       int     `json:"account_id"`
	OperationTypeId int     `json:"operation_type_id"`
	Amount          float64 `json:"amount"`
}

func ParseTransaction(bodyReader io.ReadCloser) Transaction {
	decoder := json.NewDecoder(bodyReader)
	var t Transaction
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	return t
}
