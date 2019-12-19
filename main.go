package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world!!!")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}
	log.Printf("[MAIN] Server is running at 0.0.0.0:%v\n", port)
	log.Println("[MAIN]", e.Start(":"+port))
}
