package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		//AllowCredentials: true,
	}))

	// Routes
	e.GET("/ping", ping)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

func ping(c echo.Context) error {
	return c.JSON(http.StatusOK, struct {
		Value string `json:"value"`
	}{
		Value: "pong!!",
	})
}
