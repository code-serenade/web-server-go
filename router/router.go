package router

import (
	"github.com/code-serenade/web-server-go/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", handler.Ping)
	return r
}
