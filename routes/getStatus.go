package routes

import (
	"github.com/chentekdev/tsconector/peticiones"
	"github.com/gin-gonic/gin"
)

func GetStatus(c *gin.Context) {
	msg, _, _ := peticiones.Status()
	c.JSON(200, gin.H{
		"exito": true,
		"msg":   msg,
	})

}
