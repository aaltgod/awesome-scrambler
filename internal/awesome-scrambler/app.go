package awesome_scrambler

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"log"

	"github.com/alyaskastorm/awesome-scrambler/internal/delivery/http/api"
	"github.com/alyaskastorm/awesome-scrambler/internal/repository"
)


func RunApp() {

	client, err := repository.CreateConnection()
	if err != nil {
		log.Fatalln(err)
	}
	client.Database("storage").Drop(context.TODO())
	client.Disconnect(context.TODO())

	db := repository.NewTextStorage()
	h := api.NewHandler(db)

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{}))

	e.POST("/api/encrypt_text", h.EncryptText)
	e.POST("/api/get_cipher_text", h.GetCipherText)

	e.Logger.Fatal(e.Start(":5000"))
}
