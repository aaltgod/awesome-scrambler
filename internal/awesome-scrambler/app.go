package awesome_scrambler

import (
	"context"
	api2 "github.com/alyaskastorm/awesome-scrambler/internal/awesome-scrambler/delivery/http/api"
	repository2 "github.com/alyaskastorm/awesome-scrambler/internal/awesome-scrambler/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"log"
)


func RunApp() {

	client, err := repository2.CreateConnection()
	if err != nil {
		log.Fatalln(err)
	}
	client.Database("storage").Drop(context.TODO())
	client.Disconnect(context.TODO())

	db := repository2.NewTextStorage()
	h := api2.NewHandler(db)

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{}))

	e.POST("/api/encrypt_text", h.EncryptText)
	e.POST("/api/get_cipher_text", h.GetCipherText)

	e.Logger.Fatal(e.Start(":5000"))
}
