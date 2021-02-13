package util

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"pismo-challenge-go/models"
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

func ParseTransaction(bodyReader io.ReadCloser) models.Transaction {
	decoder := json.NewDecoder(bodyReader)
	var t models.Transaction
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	return t
}
