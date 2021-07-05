package awesome_scrambler

import (
	"context"
	"github.com/alyaskastorm/awesome-scrambler/internal/awesome-scrambler/delivery/http/api"
	storage "github.com/alyaskastorm/awesome-scrambler/internal/awesome-scrambler/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoLog "github.com/labstack/gommon/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"log"
)


func RunApp() {

	client, err := storage.CreateConnection()
	if err != nil {
		log.Fatalln(err)
	}
	client.Database("storage").Drop(context.TODO())
	client.Disconnect(context.TODO())

	db := storage.NewTextStorage()
	h := api.NewHandler(db)

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{}))

	e.Logger.SetLevel(echoLog.INFO)

	e.POST("/api/encrypt_text", h.EncryptText)
	e.POST("/api/get_cipher_text", h.GetCipherText)

	go func(){
		if err := e.Start(":5000"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("The service is shutting down")
		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	killSignal := <-interrupt
	switch killSignal {
	case os.Interrupt:
		log.Println("Got SIGINT")
	case syscall.SIGTERM:
		log.Println("Got SIGTERM")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
