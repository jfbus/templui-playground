//go:generate templplaygroundgenerator
package main

import (
	"fmt"
	"os"

	"github.com/jfbus/templ-components/cmd/templplayground/assets"
	"github.com/jfbus/templ-components/cmd/templplayground/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		_, _ = fmt.Fprintln(os.Stderr, "PORT must be set") // nolint
		os.Exit(1)
	}

	e := echo.New()

	e.StaticFS("/static", assets.Assets)
	e.GET("/", handlers.Index)
	e.POST("/view/:component", handlers.ViewComponent)
	if err := e.Start(":" + port); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err) // nolint
		os.Exit(1)
	}

}
