package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	// htmlのディレクトリを指定
    engine.LoadHTMLGlob("templates/*")
	ua := ""
     // ミドルウェアを使用
    engine.Use(func(c *gin.Context) {
        ua = c.GetHeader("User-Agent")
        c.Next()
    })
	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello GORIRA",
			"User-Agent": ua,
		})
	})
	engine.GET("/template", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", gin.H{
             // htmlに渡す変数を定義
            "message": "hello gin",
        })
    })
	engine.Run(":3000")
}
