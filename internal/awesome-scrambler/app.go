package awesome_scrambler

import (
	"github.com/alyaskastorm/awesome-scrambler/internal/delivery/http/api"
	"github.com/alyaskastorm/awesome-scrambler/internal/repository"
	"github.com/labstack/echo/v4"
	"log"
)


func RunApp() {

	if err := repository.CreateConnection(); err != nil {
		log.Fatalln(err)
	}

	h := api.NewHandler("db conn")

	e := echo.New()


	e.POST("/api/encrypt_text", h.EncryptText)
	e.POST("/api/get_cipher_text", h.GetCipherText)

	e.Logger.Fatal(e.Start(":5000"))
}
