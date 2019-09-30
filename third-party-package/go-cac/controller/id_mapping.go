package controller

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/patrickmn/go-cache"
	"net/http"
	"time"
)

type (
	IDMapping struct {
		MID string `json:"mid"`
		OID string `json:"oid"`
	}

	PostIDMappingBody struct {
		Result int `json:"result"`
	}
)

var (
	// 10秒間キャッシュ、1秒ごとに期限切れのキャッシュをクリーン
	ca = cache.New(10*time.Second, 1*time.Second)
)

/**
 * IDマッピング取得
 */
func GetIDMapping(c echo.Context) error {
	mid := c.Param("mid")

	oid, found := ca.Get(mid)
	if !found {
		fmt.Println("mid: %s", mid)
		return c.JSON(http.StatusNotFound, "not found oid to cache")
	}
	fmt.Println("mid: %s oid: %s", mid, oid)
	return c.JSON(http.StatusOK, IDMapping{MID: mid, OID: oid.(string)})
}

/**
 * IDマッピングキャッシュ
 */
func PostIDMapping(c echo.Context) error {
	IDMappingRequest := new(IDMapping)
	if err := c.Bind(IDMappingRequest); err != nil {
		fmt.Println("IDMappingRequest: %V", IDMappingRequest)
		return c.JSON(http.StatusInternalServerError, "bind error")
	}

	ca.Set(IDMappingRequest.MID, IDMappingRequest.OID, cache.DefaultExpiration)
	fmt.Println("mid: %s oid: %s", IDMappingRequest.MID, IDMappingRequest.OID)

	return c.JSON(http.StatusOK, PostIDMappingBody{Result: 200})
}
