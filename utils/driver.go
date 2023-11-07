package utils

import (
	"context"
	"fmt"

	"github.com/roby-aw/go-clean-architecture-hexagonal/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DatabaseDriver string

const (
	MongoDB DatabaseDriver = "mongodb"
)

type DatabaseConnection struct {
	Driver DatabaseDriver

	MongoDB     *mongo.Database
	mongoClient *mongo.Client
}

func NewConnectionDatabase(config *config.AppConfig) *DatabaseConnection {
	var db DatabaseConnection
	db.mongoClient = newMongodb(config)
	db.MongoDB = db.mongoClient.Database(config.Database.DBNAME)

	return &db
}

func newMongodb(config *config.AppConfig) *mongo.Client {
	url := fmt.Sprintf("mongodb://%s:%s@%s:%s/?authSource=%s&authMechanism=SCRAM-SHA-256", config.Database.DBUSER, config.Database.DBPASS, config.Database.DBURL, config.Database.DBPORT, config.Database.DBNAME)

	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		panic(err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		panic(err)
	}
	return client
}

func (db *DatabaseConnection) CloseConnection() {
	db.mongoClient.Disconnect(context.Background())
}
