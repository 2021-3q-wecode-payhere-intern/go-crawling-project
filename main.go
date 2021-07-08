package main

import (
	"codef"
	"context"
	"db"
	"log"
)

func main() {
	// db커넥션
	client, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Collection("testStore")

	codefResult := codef.GetDepositInfos("20210501", "20210505")

	for _, tempMap := range codefResult {
		value := struct {
			DommEndDate          string
			CommMemberStoreGroup string
			CommStartDate        string
			ResAccountIn         string
			ResBankName          string
			ResCardCompany       string
			ResDepositDate       string
			ResMemberStoreNo     string
			ResOtherDeposit      string
			ResPaymentAccount    string
			ResSalesAmount       string
			ResSalesCount        string
			ResSuspenseAmount    string
		}{
			tempMap["commEndDate"],
			tempMap["commMemberStoreGroup"],
			tempMap["commStartDate"],
			tempMap["resAccountIn"],
			tempMap["resBankName"],
			tempMap["resCardCompany"],
			tempMap["resDepositDate"],
			tempMap["resMemberStoreNo"],
			tempMap["resOtherDeposit"],
			tempMap["resPaymentAccount"],
			tempMap["resSalesAmount"],
			tempMap["resSalesCount"],
			tempMap["resSuspenseAmount"],
		}

		_, err := collection.InsertOne(context.Background(), value)
		if err != nil {
			log.Fatal(err)
		}
	}

	// db 커넥션 종료
	client.Client().Disconnect(context.TODO())
}
