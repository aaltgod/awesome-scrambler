package email_gmail

import (
	"context"
	"encoding/base64"
	"fmt"
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
}

func NewGmailService(gs *gmail.Service) *GmailService {
	return &GmailService{Service: gs}
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

	unreadMessageBody := message.Snippet

	cipherText, key, err := encrypter.Encrypt(unreadMessageBody)
	if err != nil {
		return false, err
	}

	emailBody := fmt.Sprintf("Key: %s\n\n%s", key, cipherText)

	var messageToSend gmail.Message

	emailTo := "To: " + sender + "\r\n"
	subject := "Subject: " + "Your cipher text\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	msg := []byte(emailTo + subject + subject + mime + "\n" + emailBody)

	messageToSend.Raw = base64.URLEncoding.EncodeToString(msg)

	_, err = gs.Service.Users.Messages.Send("me", &messageToSend).Do()
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

	if subject == "Encrypt" {
		return true
	}

	return false
}

func RunApp() {

	var (
		gmailService = NewGmailService(OAuthGmailService())
		wg = &sync.WaitGroup{}
	)

	for {
		time.Sleep(10 * time.Second)

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