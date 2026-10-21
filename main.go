package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{
			"Title": "Go Public Board",
		})
	})
	r.GET("/a200", func(ctx *gin.Context) {
		ctx.HTML(200, "a200.html", gin.H{
			"Title": "Go h Board",
		})
	})
	r.Static("/public", "./public")

	r.Run(":8080")
}
