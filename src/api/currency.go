package api

import (
	"crud-api/config"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
)

func CurrencySort(c *gin.Context) {
	var (
		data     []config.Request
		currency []string
	)
	for i := 0; i < len(config.RequestData); i++ {
		currency = append(currency, config.RequestData[i].Currency)
	}
	s := Unique(currency)
	sort.Strings(s)
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(currency); j++ {
			if s[i] == config.RequestData[j].Currency {
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
