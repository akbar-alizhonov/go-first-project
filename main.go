package main

import (
	"yandex_go/cmd/handlers"
	"yandex_go/cmd/storage"

	"github.com/labstack/echo/v4"
)



func main() {
	e := echo.New()
	storage.InitDB()
	e.GET("/", handlers.Home)
	e.POST("/users", handlers.CreateUser)
	e.Logger.Fatal(e.Start(":8000"))
}