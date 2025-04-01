package main

import (
	"github.com/ShudderStorm/go-github-tracker/cmd/app/gin"
	"github.com/ShudderStorm/go-github-tracker/cmd/app/internal/oauth"
)

func main() {
	oauth.SetRedirectUrl("http://localhost:8080/callback")
	gin.Run(gin.DefaultPort)
}
