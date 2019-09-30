package router

import (
	"../controller"
	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()

	// Routes
	v1 := e.Group("/cache")
	{
		// キャッシュ
		v1.GET("/id-mapping/:mid", controller.GetIDMapping)
		v1.POST("/id-mapping", controller.PostIDMapping)
	}
	return e
}
