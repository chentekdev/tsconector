package main

import (
	"fmt"

	"github.com/chentekdev/tsconector/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Bienvenido al conector Tecnosinergia")

	r := gin.Default()
	r.GET("/status", gin.HandlerFunc(routes.GetStatus))
	r.GET("/getinventario", gin.HandlerFunc(routes.GetInventario))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "exito",
		})
	})

	r.Run()
}
