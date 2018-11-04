package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("---from8999to8888---")

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hellllllo")

	})

	r.POST("/", PostRouteHandler)


	r.Run(":4567")
}

func PostRouteHandler(c *gin.Context) {
	c.Request.ParseForm()
	fmt.Println("crawl target URL")
	fmt.Println(c.Request.Form["targetURL"])

	c.String(200, c.Request.Form["targetURL"][0])
}
