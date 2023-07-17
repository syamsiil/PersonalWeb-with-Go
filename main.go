package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
    e := echo.New()

    // e = echo package
	// GET =  run the method
	// "/" = endpoint/routing ("localhost:5000 , ex. "/home")
	// helloWorld = function that will run if the route are opened
    e.GET("/", helloWorld)
 
    e.Logger.Fatal(e.Start("localhost:5000"))
}

func helloWorld(c echo.Context)error {
	return c.String(http.StatusOK, "Hello World")
}