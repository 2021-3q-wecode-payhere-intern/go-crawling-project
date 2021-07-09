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
		MongoDBHost:     config.MongoDBHost,
		MongoDBPort:     config.MongoDBPort,
		MongoDBName:     config.MongoDBName,
		MongoDBUserName: config.MongoDBUserName,
		MongoDBPassword: config.MongoDBPassword,
	}

	//Codef instance
	codeF := codef.CodefConfig{
		CodefPublicKey:    config.CodefPublicKey,
		CodefClientId:     config.CodefClientId,
		CodefClientSecret: config.CodefClientSecret,
		CrefiaId:          config.CrefiaId,
		CrefiaPassword:    config.CrefiaPassword,
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
