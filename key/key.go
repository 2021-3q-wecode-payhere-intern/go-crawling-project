package key

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type keyDatas struct {
	Codef, Database map[string]string
}

func getFile() []byte {
	byteValue, err := ioutil.ReadFile("./key/key.json")
	if err != nil {
		log.Fatal(err)
	}

	return byteValue
}

func GetCodefKeyValue() map[string]string {
	var key keyDatas

	byteValue := getFile()

	json.Unmarshal(byteValue, &key)

	return key.Codef
}

func GetDatabaseKeyValue() map[string]string {
	var key keyDatas

	byteValue := getFile()

	json.Unmarshal(byteValue, &key)

	return key.Database
}
