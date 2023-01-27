package api

import (
	"crud-api/config"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Filter(c *gin.Context) {

	key, _ := c.GetQuery("key")
	value, _ := c.GetQuery("name")

	if key == "" || value == "" {

		c.JSON(400, gin.H{
			"error":   true,
			"message": "mismatch or empty query",
		})
		return
	}
	var (
		data    []config.Request
		keydata []string
	)
	for i := 0; i < len(config.RequestData); i++ {
		if key == "country" {
			keydata = append(keydata, config.RequestData[i].Country)
		} else if key == "continent" {
			keydata = append(keydata, config.RequestData[i].Continent)
		} else if key == "currency" {
			keydata = append(keydata, config.RequestData[i].Currency)
		}
		if strings.Contains(strings.ToLower(keydata[i]), strings.ToLower(value)) {
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
