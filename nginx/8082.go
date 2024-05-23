package main

import (
	"github.com/gin-gonic/gin"
	"math/rand/v2"
	"net/http"
)

func main() {
	r := gin.Default()
	// r.Use(Cors())
	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		v1.GET("/ccc", func(c *gin.Context) {
			rand.Int64()
			c.JSON(200, gin.H{
				"whoami":  "8082",
				"message": rand.Int64(),
			})
		})
	}
	r.Run(":8082")
}

// Cors 跨域中间件
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置允许跨域的域名，*代表允许任意域名跨域
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		// 设置允许的请求方法，*代表允许任意方法，如GET/POST/PUT等
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		// 设置允许的请求头
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		// 设置允许带上cookie
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法，因为有的浏览器的预请求会先发一个OPTIONS请求
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		// 处理请求
		c.Next()
	}
}
