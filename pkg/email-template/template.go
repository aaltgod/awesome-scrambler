package email_template

import (
	"fmt"
	"github.com/matcornic/hermes/v2"
	"log"
)

func GenerateTemplate(to, key, cipherText string) string {

	h := hermes.Hermes{
		Product: hermes.Product{
			Name: "Awesome scrambler",
		},
	}

	email := hermes.Email{
		Body: hermes.Body{
			Name: to,
			Greeting: "Hello, here is your ciphertext.",
			FreeMarkdown: hermes.Markdown(
				fmt.Sprintf("*Key*: %s\n\n*Ciphertext*: %s", key, cipherText),
				),
			Outros: []string{"Goodbye"},
		},
	}

	res, err := h.GeneratePlainText(email)
	if err != nil {
		log.Println(err)
	}

	return res
}