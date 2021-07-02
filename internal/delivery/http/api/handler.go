package api

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type Handler struct {
	db string
}

type Text struct {
	Text string `json:"text"`
}

type CipherText struct {
	Key string `json:"key,omitempty"`
	CipherText string `json:"cipher_text,omitempty"`
	Link string `json:"link,omitempty"`
}

func NewHandler(db string) *Handler {
	return &Handler{db: "Data base"}
}

func (h *Handler) EncryptText(c echo.Context) error {
	var text Text

	if err := c.Bind(&text); err != nil {
		log.Println("[EncryptText-BIND]: ", err)
		return err
	}

	response := &CipherText{
		Key: "qweqwe123",
		Link: "234234234sfadkgjhsdfg",
	}

	return c.JSON(http.StatusAccepted, response)
}

func (h *Handler) GetCipherText(c echo.Context) error {
	var link CipherText

	if err := c.Bind(&link); err != nil {
		log.Println("[GetCipherText-BIND]: ", err)
		return err
	}

	cipherText := "sdsd"
	response := &CipherText{
		CipherText: cipherText,
	}

	return c.JSON(http.StatusAccepted, response)
}


