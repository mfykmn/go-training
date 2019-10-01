package main

import (
	"github.com/labstack/echo/engine/fasthttp"
	"./router"
)

func main() {
	route := router.Init()
	route.Run(fasthttp.New(":60001"))
}