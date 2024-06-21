package main

import (
	"context"
	"log"
	"pair_programming/config"
	"pair_programming/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	client, db := config.Connect(context.Background(), "pair_programming")

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	router.InitRoutes(e, db)

	e.Logger.Fatal(e.Start(":8080"))
}
