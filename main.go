package main

import (
	"codef"
	"config"
	"context"
	"db"
	"log"
)

func main() {
	// init config
	config := config.InitConfig()

	// MongoDB instance
	mongoDB := db.MongoDBConfig{
		config.MongoDBHost,
		config.MongoDBPort,
		config.MongoDBName,
		config.MongoDBUserName,
		config.MongoDBPassword,
	}

	//Codef instance
	codeF := codef.CodefConfig{
		config.CodefPublicKey,
		config.CodefClientId,
		config.CodefClientSecret,
		config.CrefiaId,
		config.CrefiaPassword,
	}

	// db커넥션
	mongoDBClient, err := mongoDB.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	mongoDBCollection := mongoDBClient.Collection("testStore")

	codefResult := codeF.GetDepositInfos("20210501", "20210505")

	for _, tempMap := range codefResult {
		value := struct {
			CommEndDate          string
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

		_, err := mongoDBCollection.InsertOne(context.Background(), value)
		if err != nil {
			log.Fatal(err)
		}
	}

	// db 커넥션 종료
	mongoDBClient.Client().Disconnect(context.TODO())
}
