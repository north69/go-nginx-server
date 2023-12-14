package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Hello")

	server := echo.New()

	server.GET("/status", Handler)

	server.Logger.Fatal(server.Start(":8081"))
}

func Handler(c echo.Context) error {
	d := time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)
	dur := time.Until(d)
	s := fmt.Sprintf("Days remained: %d", int64(dur.Hours()/24))
	return c.String(http.StatusOK, s)
}
