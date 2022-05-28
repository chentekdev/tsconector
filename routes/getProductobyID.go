package routes

import (
	"github.com/chentekdev/tsconector/peticiones"
	"github.com/gin-gonic/gin"
)

func GetProductobyID(c *gin.Context) {
	type ResID struct {
		ID string `json:"id"`
	}

	res := ResID{}
	err := c.BindJSON(&res)
	if err != nil {
		c.JSON(400, gin.H{
			"exito": false,
			"msg":   "Error: Datos invalidos " + err.Error(),
		})
		return
	}

	p, msg, exito, err := peticiones.GetProductobyID(res.ID)

	if !exito {
		c.JSON(400, gin.H{
			"exito": false,
			"msg":   msg + " " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"exito": true,
		"msg":   p,
	})

}
