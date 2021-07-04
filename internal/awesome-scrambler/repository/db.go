package repository

import (
	"context"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CipherText struct {
	Link string `bson:"link"`
	CipherText string `bson:"cipher_text"`
}

type Storage interface {
	InsertText(cipherText, link string) error
	GetCipherText(link string) (string, error)
}

type TextStorage struct {
	mu *sync.Mutex
}

func NewTextStorage() *TextStorage {
	return &TextStorage{mu: &sync.Mutex{}}
}

func CreateConnection() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017/storage")

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

func (ts *TextStorage) InsertText(cipherText, link string) error {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	client, err := CreateConnection()
	if err != nil {
		return err
	}

	collection := client.Database("storage").Collection("text")
	_, err = collection.InsertOne(context.TODO(), CipherText{
		CipherText: cipherText,
		Link: link,
	})
	if err != nil {
		return err
	}
	defer client.Disconnect(context.TODO())


	return nil
}

func (ts *TextStorage) GetCipherText(link string) (string, error) {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	client, err := CreateConnection()
	if err != nil {
		return "", err
	}

	var result CipherText

	collection := client.Database("storage").Collection("text")
	err = collection.FindOne(context.TODO(), bson.D{{"link", link}}).Decode(&result)
	if err != nil {
		log.Println("[FIND-ONE]: ", err)
		return "", err
	}
	defer client.Disconnect(context.TODO())

	return result.CipherText, nil
}
