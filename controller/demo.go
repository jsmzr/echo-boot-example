package controller

import (
	"net/http"

	log "github.com/jsmzr/boot-log"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func Get(c echo.Context) error {
	log.Info("request /demo")
	return c.String(http.StatusOK, "hello world!")
}

func GetConfig(c echo.Context) error {
	key := c.QueryParam("key")
	return c.String(http.StatusOK, viper.GetString(key))
}
