package main

import (
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
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

//go:embed *.html
//go:embed rest_api_example.png
var f embed.FS

func main() {
	if mode == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.Use(gin.Recovery())
	// https://github.com/gin-gonic/examples/blob/master/assets-in-binary/example02/main.go
	tmpl := template.Must(template.New("").ParseFS(f, "*.html"))
	r.SetHTMLTemplate(tmpl)

	r.GET("/", indexHandler)
	r.GET("/ws", wsHandler)
	r.GET("/api", apiHandler)
	r.GET("/rest_api_example.png", staticFileHandler)
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{"code": 200})
	})

	fmt.Println(r.Run(":8888"))
}
