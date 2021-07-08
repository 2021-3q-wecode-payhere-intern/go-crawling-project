package codef

import (
	"encoding/json"
	"key"
	"log"
	"strings"

	ecg "github.com/codef-io/easycodefgo"
)

type depositDatas struct {
	Data   []map[string]string
	Result map[string]string
}

func GetDepositInfos(start, end string) []map[string]string {
	startDate := strings.ReplaceAll(start, " ", "")
	endDate := strings.ReplaceAll(end, " ", "")

	if startDate == "" || endDate == "" {
		log.Fatal("invalid_date")
	}

	// codef key값
	codefMap := key.GetCodefKeyValue()
	publicKey := codefMap["public_key"]
	clientId := codefMap["client_id"]
	clientSecret := codefMap["client_secret"]
	id := codefMap["id"]
	password := codefMap["password"]

	// 코드에프 인스턴스 생성
	codefApi := &ecg.Codef{
		PublicKey: publicKey,
	}

	codefApi.SetClientInfoForDemo(clientId, clientSecret)

	pwd, err := ecg.EncryptRSA(password, publicKey)
	if err != nil {
		log.Fatal(err)
	}

	parameter := map[string]interface{}{
		"organization":     "0323",
		"id":               id,
		"password":         pwd,
		"startDate":        startDate,
		"endDate":          endDate,
		"memberStoreGroup": "",
	}

	productURL := "/v1/kr/card/a/cardsales/deposit-list"

	// 서비스타입
	// 0 : ecg.TypeProduct
	// 1 : ecg.TypeDemo
	// 2 : ecg.TypeSandbox
	codefResult, err := codefApi.RequestProduct(productURL, ecg.TypeDemo, parameter)
	if err != nil {
		log.Fatal(err)
	}

	// codef request 결과 map으로 변환
	var datas depositDatas
	json.Unmarshal([]byte(codefResult), &datas)

	if datas.Result["code"] != "CF-00000" {
		errorMsg := "codef_error : " +
			datas.Result["code"] +
			datas.Result["extraMessage"]
		log.Fatal(errorMsg)
	}

	return datas.Data
}
