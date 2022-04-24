package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Rewphg/iambot/src/logger/debug"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	if err := debug.InitDebugLogger(); err != nil {
		log.Fatal(err)
	}
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	go func() {
		port := os.Getenv("HTTP_PORT")
		if port == "" {
			port = "8000"
		}
		if err := e.Start(":" + port); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatalf("crashed: %v\n", err)
			log.Panicln("server crashed")
		}
	}()

	e.HTTPErrorHandler = customHTTPErrorHandler
	// e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
	// 	return func(c echo.Context) error {
	// 		// Extract the credentials from HTTP request header and perform a security
	// 		// check

	// 		// For invalid credentials
	// 		return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")

	// 		// For valid credentials call next
	// 		// return next(c)
	// 	}
	// })

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	// block until received interrupt signal
	<-quit

	log.Println("starting server shutdown")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		log.Panic("failed to shutdown server properly. check server logs for more info")
		e.Logger.Fatalf("failed to shutdown: %v\n", err)
	} else {
		log.Println("server shutdown down properly")
	}

}

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	errorPage := fmt.Sprintf("%d.html", code)
	if err := c.File(errorPage); err != nil {
		c.Logger().Error(err)
	}
	c.Logger().Error(err)
}
