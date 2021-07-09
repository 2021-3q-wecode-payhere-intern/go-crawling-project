package codef

import (
	"encoding/json"
	"log"
	"strings"

	ecg "github.com/codef-io/easycodefgo"
)

const DEPOSIT_END_POINT = "/v1/kr/card/a/cardsales/deposit-list"
const CODEF_SUCCESS_CODE = "CF-00000"

type CodefConfig struct {
	CodefPublicKey    string
	CodefClientId     string
	CodefClientSecret string
	CrefiaId          string
	CrefiaPassword    string
}

type DepositDatas struct {
	Data   []map[string]string
	Result map[string]string
}

func (config CodefConfig) GetDepositInfos(start, end string) []map[string]string {
	startDate := strings.ReplaceAll(start, " ", "")
	endDate := strings.ReplaceAll(end, " ", "")

	if startDate == "" || endDate == "" {
		log.Fatal("invalid_date")
	}

	// 코드에프 인스턴스 생성
	codefApi := &ecg.Codef{
		PublicKey: config.CodefPublicKey,
	}

	codefApi.SetClientInfoForDemo(config.CodefClientId, config.CodefClientSecret)

	pwd, err := ecg.EncryptRSA(config.CrefiaPassword, config.CodefPublicKey)
	if err != nil {
		log.Fatal(err)
	}

	parameter := map[string]interface{}{
		"organization":     "0323",
		"id":               config.CrefiaId,
		"password":         pwd,
		"startDate":        startDate,
		"endDate":          endDate,
		"memberStoreGroup": "",
	}

	// 서비스타입
	// 0 : ecg.TypeProduct
	// 1 : ecg.TypeDemo
	// 2 : ecg.TypeSandbox
	codefResult, err := codefApi.RequestProduct(DEPOSIT_END_POINT, ecg.TypeDemo, parameter)
	if err != nil {
		log.Fatal(err)
	}

	// codef request 결과 map으로 변환
	var datas DepositDatas
	json.Unmarshal([]byte(codefResult), &datas)

	if datas.Result["code"] != CODEF_SUCCESS_CODE {
		errorMsg := "codef_error : " +
			datas.Result["code"] +
			datas.Result["extraMessage"]
		log.Fatal(errorMsg)
	}

	return datas.Data
}
