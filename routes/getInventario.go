package routes

import (
	"github.com/chentekdev/tsconector/peticiones"
	"github.com/gin-gonic/gin"
)

func GetInventario(c *gin.Context) {
	/*type Request struct {
		Marca string `json:"marca"`
	}

	req := Request{}

	err := c.BindJSON(&req)

	if err != nil {
		c.JSON(200, gin.H{
			"exito": false,
			"msg":   "Fallo",
		})
		return
	}
	*/
	msg, _, _ := peticiones.Inventario()
	c.JSON(200, gin.H{
		"exito": true,
		"msg":   msg,
	})

}
