package api

import (
	repository2 "github.com/alyaskastorm/awesome-scrambler/internal/awesome-scrambler/repository"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"

	"github.com/alyaskastorm/awesome-scrambler/pkg/encrypter"
	"github.com/alyaskastorm/awesome-scrambler/pkg/random-string"
)

type Handler struct {
	db repository2.Storage
}

type Text struct {
	Text string `json:"text"`
}

type CipherText struct {
	Key string `json:"key,omitempty"`
	CipherText string `json:"cipher_text,omitempty"`
	Link string `json:"link,omitempty"`
}

func NewHandler(db repository2.Storage) *Handler {
	return &Handler{db: db}
}

func (h *Handler) EncryptText(c echo.Context) error {
	var text Text

	if err := c.Bind(&text); err != nil {
		log.Println("[EncryptText-BIND]: ", err)
		return err
	}

	link := random_string.GetRandomString(6)

	cipherText, key, err := encrypter.Encrypt(text.Text)
	if err != nil {
		return err
	}

	if err = h.db.InsertText(cipherText, link); err != nil {
		return err
	}

	response := &CipherText{
		Key: key,
		Link: link,
	}

	return c.JSON(http.StatusAccepted, response)
}

func (h *Handler) GetCipherText(c echo.Context) error {
	var link CipherText

	if err := c.Bind(&link); err != nil {
		log.Println("[GetCipherText-BIND]: ", err)
		return err
	}

	cipherText, err := h.db.GetCipherText(link.Link)
	if err != nil {
		log.Println("[GetCipherText-CipherText]: ", err)
		return err
	}

	response := &CipherText{
		CipherText: cipherText,
	}

	return c.JSON(http.StatusAccepted, response)
}


