package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.Static("/", "view")

	r.Run(":8000")
}
