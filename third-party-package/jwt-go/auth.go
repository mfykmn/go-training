package main

import (
	"net/http"
	"os"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
)

// GetTokenHandler get token
func GetTokenHandler(rw http.ResponseWriter, r *http.Request) {
	// TODO 認証チェック


	// headerのセット
	token := jwt.New(jwt.SigningMethodHS256)

	// claimsのセット
	claims := token.Claims.(jwt.MapClaims)
	claims["admin"] = true
	claims["iss"] = "mafuyuk.com"
	claims["sub"] = "54546557354"
	claims["name"] = "taro"
	claims["iat"] = time.Now() // JWTを発行した時刻
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // 有効期限

	// 電子署名
	tokenString, _ := token.SignedString([]byte(os.Getenv("SIGNINGKEY")))

	// JWTを返却
	rw.Write([]byte(tokenString))
}

// JwtMiddleware check token
var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SIGNINGKEY")), nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})
