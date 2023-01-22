package api

import (
	"crud-api/config"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Filter(c *gin.Context) {
	countryname, _ := c.GetQuery("country-name")
	conitinent, _ := c.GetQuery("continent-name")
	currency, _ := c.GetQuery("currency-name")

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
	}

	if conitinent != "" {
		for i := 0; i < len(config.RequestData); i++ {
			if strings.Contains(strings.ToLower(config.RequestData[i].Continent), strings.ToLower(conitinent)) {
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
	}

	if currency != "" {
		for i := 0; i < len(config.RequestData); i++ {
			if strings.Contains(strings.ToLower(config.RequestData[i].Currency), strings.ToLower(currency)) {
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
	}
}
