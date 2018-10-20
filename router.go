package main

import (
	"encoding/base64"
	"net/http"
	"strings"
	"zyjsxy-api/database"
	"zyjsxy-api/util/aes"

	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/api/user/info", system.GetSelfInfoApi)
}
