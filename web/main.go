package main

import (
	"github.com/gin-gonic/gin"
	"web/controller"
)

func main() {
	r := gin.Default()

	r.Static("/tenement", "view")

	r.GET("/api/v1.0/session", controller.GetSession)

	r.Run(":8000")
}
