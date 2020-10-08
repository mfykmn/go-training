package router

import (
	"net/http"

	"github.com/mfykmn/go-training/third-party-package/go-cac/controller"

	"github.com/webx-top/echo"
)

func Init() *echo.Echo {
	e := echo.New()

	// Routes
	v1 := e.Group("/cache")
	{
		// キャッシュ
		v1.Add(http.MethodGet, "/id-mapping/:mid", controller.GetIDMapping)
		v1.Add(http.MethodPost, "/id-mapping", controller.PostIDMapping)
	}
	return e
}
