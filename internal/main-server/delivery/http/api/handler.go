package api

import (
	"log"
	"net/http"

	repos "github.com/aaltgod/awesome-scrambler/internal/main-server/repository"
	"github.com/labstack/echo/v4"

	"github.com/aaltgod/awesome-scrambler/pkg/encrypter"
)

type Handler struct {
	db repos.Storage
}

type Text struct {
	Text string `json:"text"`
}

type CipherText struct {
	Key        string `json:"key,omitempty"`
	CipherText string `json:"cipher_text,omitempty"`
}

func NewHandler(db repos.Storage) *Handler {
	return &Handler{db: db}
}

func (h *Handler) EncryptText(c echo.Context) error {
	var text Text

	if err := c.Bind(&text); err != nil {
		log.Println("[EncryptText-BIND]: ", err)
		return err
	}

	cipherText, key, err := encrypter.Encrypt(text.Text)
	if err != nil {
		return err
	}

	if err = h.db.InsertText(cipherText, key); err != nil {
		return err
	}

	response := &CipherText{
		Key: key,
	}

	return c.JSON(http.StatusAccepted, response)
}

func (h *Handler) GetCipherText(c echo.Context) error {
	var request CipherText

	if err := c.Bind(&request); err != nil {
		log.Println("[GetCipherText-BIND]: ", err)
		return err
	}

	cipherText, err := h.db.GetCipherText(request.Key)
	if err != nil {
		log.Println("[GetCipherText-CipherText]: ", err)
		return err
	}

	response := &CipherText{
		CipherText: cipherText,
	}

	return c.JSON(http.StatusAccepted, response)
}
