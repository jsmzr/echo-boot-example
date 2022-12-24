package router

import (
	"github.com/jsmzr/boot-echo"
	"github.com/jsmzr/echo-boot-example/controller"
	"github.com/labstack/echo/v4"
)

func initDemoRouter(e *echo.Echo) {
	e.GET("/demo", controller.Get)
	e.GET("/demo/config", controller.GetConfig)
}

func init() {
	boot.RegisterRouter(initDemoRouter)
}
