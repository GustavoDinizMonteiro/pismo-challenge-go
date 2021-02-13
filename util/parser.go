package util

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

func ParseBody(bodyReader io.ReadCloser) map[string]string {
	body, err := ioutil.ReadAll(bodyReader)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)

	return keyVal
}
