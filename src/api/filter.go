package api

import (
	"crud-api/config"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Filter(c *gin.Context) {
	countryname, _ := c.GetQuery("country-name")
	var data []config.Request
	if countryname != "" {
		for i := 0; i < len(config.RequestData); i++ {
			if strings.Contains(strings.ToLower(config.RequestData[i].Country), strings.ToLower(countryname)) {
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
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"error":   false,
			"message": "country name is required",
		})
		return
	}

}
