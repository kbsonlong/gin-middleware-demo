package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func MdOne() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("开始执行第一个 gin 中间件:" + time.Now().String())
		c.Next()

		fmt.Println("第一个 gin 中间件返回内容:" + time.Now().String())
	}
}

func MdTwo() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("开始执行第二个 gin 中间件:" + time.Now().String())
		c.Next()

		fmt.Println("第二个 gin 中间件返回内容:" + time.Now().String())
	}
}

func MdThree() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("开始执行第三个 gin 中间件:" + time.Now().String())

		// c.Next()
		fmt.Println("第三个 gin 中间件返回内容:" + time.Now().String())
	}
}
