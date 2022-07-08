package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/kbsonlong/gin-middleware-demo/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleware.MdOne())

	r.Use(middleware.MdThree())
	r.Use(middleware.MdTwo())
	gin.SetMode("debug")

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "test",
		})
	})
	return r
}
