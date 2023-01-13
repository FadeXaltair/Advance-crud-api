package api

import (
	"crud-api/config"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Details(c *gin.Context) {
	id, _ := c.GetQuery("id")
	userid, _ := strconv.Atoi(id)
	var data []config.Request

	for i := 0; i < len(config.RequestData); i++ {
		if userid == config.RequestData[i].Id {
			data = append(data, config.Request{
				Id:        config.RequestData[i].Id,
				Country:   config.RequestData[i].Country,
				Continent: config.RequestData[i].Continent,
				Currency:  config.RequestData[i].Currency,
			})
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  data,
	})
}
