package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type CipherText struct {
	Link string `bson:"link"`
	CipherText string `bson:"cipher_text"`
}

func CreateConnection() error {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:49153")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("[MONGO-CONNECT]: ", err)
		return err
	}

	if err = client.Ping(context.TODO(), nil); err != nil {
		log.Println("[PING]: ", err)
		return err
	}

	return nil
}
