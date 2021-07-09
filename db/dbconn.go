package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDBConfig struct {
	MongoDBHost     string
	MongoDBPort     string
	MongoDBName     string
	MongoDBUserName string
	MongoDBPassword string
}

func (config MongoDBConfig) ConnectDB() (*mongo.Database, error) {
	MongoDBUrl := "mongodb://" + config.MongoDBHost + ":" + config.MongoDBPort

	credential := options.Credential{
		Username: config.MongoDBUserName,
		Password: config.MongoDBPassword,
	}

	clientOptions := options.Client().ApplyURI(MongoDBUrl).SetAuth(credential)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	database := client.Database(config.MongoDBName)

	return database, err
}
