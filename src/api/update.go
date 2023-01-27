package api

import (
	"crud-api/config"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	id, _ := c.GetQuery("id")
	if id == "" {
		c.JSON(400, gin.H{
			"error":   true,
			"message": "empty or mismatch qyery params",
		})
		return
	}
	var body config.Request
	c.Bind(&body)
	dataid, _ := strconv.Atoi(id)
	for i := 0; i < len(config.RequestData); i++ {
		if dataid == config.RequestData[i].Id {
			data := config.Request{
				Id:        dataid,
				Country:   body.Country,
				Continent: body.Continent,
				Currency:  body.Currency,
			}
			config.RequestData[i] = data
			c.JSON(http.StatusOK, gin.H{
				"error":   false,
				"message": "updated successfully",
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"error":   true,
		"message": "no data updated",
	})
}
