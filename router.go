package main

import (
	"spider/rule/gszfcg"

	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/api/gszfcg/get", gszfcg.GetInfoApi)
}
