package main

import (
	"github.com/mfykmn/go-training/third-party-package/go-cac/router"

	"github.com/webx-top/echo/engine/fasthttp"
)

func main() {
	route := router.Init()
	route.Run(fasthttp.New(":60001"))
}
