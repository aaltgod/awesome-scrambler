package email_gmail

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/alyaskastorm/awesome-scrambler/internal/emai-gmail/repository"
	storage "github.com/alyaskastorm/awesome-scrambler/internal/main-server/repository"
	template "github.com/alyaskastorm/awesome-scrambler/pkg/email-templator"
	"github.com/alyaskastorm/awesome-scrambler/pkg/encrypter"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
	"log"
	"os"
	"sync"
	"time"
)

type GmailService struct {
	Service *gmail.Service
	Storage *repository.TextStorage
}

func NewGmailService(gs *gmail.Service, storage *repository.TextStorage) *GmailService {
	return &GmailService{
		Service: gs,
		Storage: storage,
	}
}

func OAuthGmailService() *gmail.Service {

	config := oauth2.Config{
		ClientID: os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Endpoint: google.Endpoint,
		RedirectURL: "http://localhost",
	}

	token := oauth2.Token{
		AccessToken: os.Getenv("ACCESS_TOKEN"),
		RefreshToken: os.Getenv("REFRESH_TOKEN"),
		TokenType: "Bearer",
		Expiry: time.Now(),
	}

	var tokenSource = config.TokenSource(context.Background(), &token)

	srv, err := gmail.NewService(
		context.Background(),
		option.WithTokenSource(tokenSource),
	)
	if err != nil {
		log.Fatalln(err)
	}

	gmailService := srv
	if gmailService != nil {
		fmt.Println("Email service is initialized")
	}

	return gmailService
}

func (gs *GmailService) GetUnreadMessages() []*gmail.Message {

	gsl, err := gs.Service.Users.Messages.List("me").Do()
	if err != nil {
		log.Println(err)
		return nil
	}

	var (
		wg             = &sync.WaitGroup{}
		unreadMessages []*gmail.Message
	)

	for _, g := range gsl.Messages {
		ID := g.Id

		wg.Add(1)
		go func(wg *sync.WaitGroup, ID string) {
			defer wg.Done()

			message, err := gs.Service.Users.Messages.Get("me", ID).Do()
			if err != nil {
				log.Fatalln(err)
			}

			for _, l := range message.LabelIds {
				if l == "UNREAD" {
					unreadMessages = append(unreadMessages, message)
				}
			}

		}(wg, ID)
	}

	wg.Wait()

	return unreadMessages
}

func (gs *GmailService) SendMessage(message *gmail.Message) (bool, error) {

	var sender string

	for _, v := range message.Payload.Headers {
		switch v.Name {
		case "Return-Path":
			sender = v.Value
			break
		}
	}

	if len(message.Payload.Parts) == 0 {
		return false, nil
	}

	body := message.Payload.Parts[0].Body.Data

	originalBodyBytes, err := base64.StdEncoding.DecodeString(body)
	if err != nil {
		return false, err
	}

	unreadMessageBody := string(originalBodyBytes)

	cipherText, key, err := encrypter.Encrypt(unreadMessageBody)
	if err != nil {
		return false, err
	}

	emailBody := template.GenerateTemplate(sender, key, cipherText)

	var messageToSend gmail.Message

	emailTo := "To: " + sender + "\r\n"
	subject := "Subject: " + "Your cipher text\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	msg := []byte(emailTo + subject + mime + "\n" + emailBody)

	messageToSend.Raw = base64.URLEncoding.EncodeToString(msg)

	_, err = gs.Service.Users.Messages.Send("me", &messageToSend).Do()
	if err != nil {
		return false, err
	}

	err = gs.Storage.InsertText(cipherText, key)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (gs *GmailService) TrashMessage(message *gmail.Message) (bool, error) {

	ID := message.Id
	_, err := gs.Service.Users.Messages.Trash("me", ID).Do()
	if err != nil {
		return false, err
	}

	return true, nil
}

func IsSubjectCorrect(message *gmail.Message) bool {

	var subject string

	LOOP:
		for _, v := range message.Payload.Headers{
			switch v.Name {
			case "Subject":
				subject = v.Value
				break LOOP
			}
		}

	return subject == "Encrypt"
}

func RunApp() {

	client, err := storage.CreateConnection()
	if err != nil {
		log.Fatalln(err)
	}
	client.Database("storage").Drop(context.TODO())
	client.Disconnect(context.TODO())

	var (
		mu = &sync.Mutex{}
		gmailService = NewGmailService(OAuthGmailService(), repository.NewTextStorage(mu))
		wg = &sync.WaitGroup{}
	)

	for {
		time.Sleep(5 * time.Second)

		unreadMessages := gmailService.GetUnreadMessages()

		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()

			for _, msg := range unreadMessages {
				if !IsSubjectCorrect(msg) {
					status, err := gmailService.TrashMessage(msg)
					if err != nil {
						log.Fatalln("TRASH", err)
					}
					if status {
						log.Println("OK TRASH NOT CORRECT SUBJECT")
					}

					continue
				}

				status, err := gmailService.SendMessage(msg)
				if err != nil {
					log.Fatalln("SEND", err)
				}
				if status {
					log.Println("OK SEND")
				}

				status, err = gmailService.TrashMessage(msg)
				if err != nil {
					log.Fatalln("TRASH", err)
				}
				if status {
					log.Println("OK TRASH")
				}
			}
		}(wg)

		wg.Wait()
	}
}