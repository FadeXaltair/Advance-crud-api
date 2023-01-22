package api

import (
	"crud-api/config"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	id, _ := c.GetQuery("id")
	dataid, _ := strconv.Atoi(id)

	index := dataid - 1

	for i := 0; i < len(config.RequestData); i++ {
		if config.RequestData[i].Id == dataid {
			config.RequestData = append(config.RequestData[:index], config.RequestData[index+1:]...)
			c.JSON(http.StatusOK, gin.H{
				"error":   false,
				"messgae": "data deleted successfully",
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"error":   true,
		"message": "no id availiable",
	})

}
