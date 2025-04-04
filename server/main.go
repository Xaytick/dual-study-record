package main

import (
    "log"
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)

func main() {
    r := gin.Default()

    // CORS 配置
    r.Use(cors.Default())

    // 测试路由
    r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Dual Study Record API is running!",
		})
	})

    // 启动服务器
    if err := r.Run(":8080"); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}