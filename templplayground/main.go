//go:generate templplaygroundgenerator
package main

import (
	"fmt"
	"os"

	"github.com/jfbus/templui-playground/templplayground/assets"
	"github.com/jfbus/templui-playground/templplayground/handlers"
	"github.com/jfbus/templui/components/style"
	"github.com/jfbus/templui/skin"
	"github.com/labstack/echo/v4"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		_, _ = fmt.Fprintln(os.Stderr, "PORT must be set") // nolint
		os.Exit(1)
	}

	style.SetSkin(skin.Default)

	e := echo.New()
	e.StaticFS("/static", assets.Assets)
	e.GET("/", handlers.Index)
	e.GET("/view/:component", handlers.ViewComponent)
	e.POST("/update/:component", handlers.UpdateComponent)
	if err := e.Start(":" + port); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err) // nolint
		os.Exit(1)
	}

}
