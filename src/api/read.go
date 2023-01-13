package api

import (
	"crud-api/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Read(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  config.RequestData,
	})

}
