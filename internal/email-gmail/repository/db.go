package repository

import (
	"context"
	"log"
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CipherText struct {
	Key        string `bson:"key"`
	CipherText string `bson:"cipher_text"`
}

type Storage interface {
	InsertText(cipherText, key string) error
}

type TextStorage struct {
	mu *sync.Mutex
}

func NewTextStorage(mu *sync.Mutex) *TextStorage {
	return &TextStorage{mu: mu}
}

func CreateConnection() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalln("[MONGO-CONNECT]: ", err)
		return &mongo.Client{}, err
	}

	if err = client.Ping(context.TODO(), nil); err != nil {
		log.Fatalln("[PING]: ", err)
		return &mongo.Client{}, err
	}

	return client, err
}

func (ts *TextStorage) InsertText(cipherText, key string) error {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	client, err := CreateConnection()
	if err != nil {
		return err
	}
	defer client.Disconnect(context.TODO())

	collection := client.Database("storage").Collection("text")
	_, err = collection.InsertOne(context.TODO(), CipherText{
		CipherText: cipherText,
		Key:        key,
	})
	if err != nil {
		return err
	}

	return nil
}
