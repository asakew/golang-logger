package main

import (
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	log.Println("http://localhost:1323")

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Use(middleware.Logger())
	// {"time":"2025-01-02T11:29:15.80439+10:00","id":"","remote_ip":"::1","host":"localhost:1323","method":"GET","uri":"/","user_agent":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36","status":200,"error":"","latency":4542,"latency_human":"4.542µs","bytes_in":0,"bytes_out":13}

	e.Logger.Fatal(e.Start(":1323"))

}
