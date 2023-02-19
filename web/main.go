package main

import (
	"github.com/gin-gonic/gin"
	"web/controller"
)

func main() {
	r := gin.Default()

	r.Static("/tenement", "view")

	v1 := r.Group("/api/v1.0")

	{
		v1.GET("/session", controller.GetSession)
		v1.GET("/image-code/:uuid", controller.GetImageCode)
	}

	r.Run(":8000")
}
