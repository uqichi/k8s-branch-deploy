package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

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

	// Set sub domain to context
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			req := c.Request()
			fmt.Println("req.Host", req.Host)
			host := strings.Split(req.Host, ":")[0]
			args := strings.Split(host, ".")
			fmt.Println("host", host)
			sub0 := args[0]
			fmt.Println("sub0", sub0)
			if sub0 == "" {
				return echo.ErrBadRequest
			}
			setCtxSub0(c, sub0)
			return next(c)
		}
	})

	// Routes
	e.GET("/ping", ping)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

func ping(c echo.Context) error {
	fmt.Println("pong!")
	fmt.Println(getCtxSub0(c))
	return c.JSON(http.StatusOK, struct {
		Value     string `json:"value"`
		SubDomain string `json:"sub_domain"`
	}{
		Value:     "pong!!",
		SubDomain: getCtxSub0(c),
	})
}

func setCtxSub0(ctx echo.Context, s string) {
	ctx.Set("ieBee3qu", s)
}

func getCtxSub0(ctx echo.Context) string {
	val := ctx.Get("ieBee3qu")
	ret, ok := val.(string)
	if !ok {
		log.Println("fail getCtxSub0")
		return ""
	}
	return ret
}
