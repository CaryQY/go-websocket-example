package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var (
	mode = "%MODE%"
)

var domain = "http://127.0.0.1:8888"

func init() {
	if mode == "prod" {
		domain = "https://gows.caryqy.top"
	}
}

func main() {
	if mode == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.Use(gin.Recovery())

	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("./*.html")
	r.GET("/", indexHandler)
	r.GET("/ws", wsHandler)
	r.GET("/api", apiHandler)
	r.GET("/rest_api_example.png", staticFileHandler)

	fmt.Println(r.Run(":8888"))
}
