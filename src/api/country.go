package api

import (
	"crud-api/config"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
)

func CountrySort(c *gin.Context) {
	var (
		data    []config.Request
		country []string
	)
	for i := 0; i < len(config.RequestData); i++ {
		country = append(country, config.RequestData[i].Country)
	}
	sort.Strings(country)
	for i := 0; i < len(country); i++ {
		for j := 0; j < len(config.RequestData); j++ {
			if country[i] == config.RequestData[j].Country {
				data = append(data, config.Request{
					Id:        config.RequestData[j].Id,
					Country:   config.RequestData[j].Country,
					Continent: config.RequestData[j].Continent,
					Currency:  config.RequestData[j].Currency,
				})
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  data,
	})
}
