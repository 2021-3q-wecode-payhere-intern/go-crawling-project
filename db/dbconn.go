package db

import (
	"context"
	"key"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectDB() (*mongo.Database, error) {
	databaseMap := key.GetDatabaseKeyValue()

	ip := databaseMap["ip"]
	port := databaseMap["port"]
	dbName := databaseMap["dbName"]
	id := databaseMap["id"]
	password := databaseMap["password"]
	dbURI := "mongodb://" + ip + ":" + port

	credential := options.Credential{
		Username: id,
		Password: password,
	}

	clientOptions := options.Client().ApplyURI(dbURI).SetAuth(credential)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	database := client.Database(dbName)

	return database, err
}
