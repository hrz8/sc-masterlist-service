package main

import (
	"fmt"
	"net/http"

	Config "github.com/hrz8/sc-masterlist-service/src/config"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	config, err := Config.NewConfig()
	if err != nil {
		fmt.Println(err)
	}
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.SERVICE_PORT)))
}
