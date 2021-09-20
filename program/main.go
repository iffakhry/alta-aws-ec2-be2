package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = ":80"
	}
	e := echo.New()
	e.GET("/", hello)
	e.GET("/users", hello)
	e.GET("/:name", helloName)
	e.Start(port)
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}

func helloName(c echo.Context) error {
	name := c.Param("name")
	return c.String(http.StatusOK, fmt.Sprintf("Hello %s", name))
}
