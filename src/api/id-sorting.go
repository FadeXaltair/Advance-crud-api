package api

import (
	"crud-api/config"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Sort(c *gin.Context) {
	firstvalue, _ := c.GetQuery("id-1")
	secondvalue, _ := c.GetQuery("id-2")
	Id1, _ := strconv.Atoi(firstvalue)
	Id2, _ := strconv.Atoi(secondvalue)
	var data1, data2 []config.Request
	if Id1 > Id2 {
		for i := Id1 - 1; i >= Id2-1; i-- {
			data1 = append(data1, config.Request{
				Id:        config.RequestData[i].Id,
				Country:   config.RequestData[i].Country,
				Continent: config.RequestData[i].Continent,
				Currency:  config.RequestData[i].Currency,
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"error": false,
			"data":  data1,
		})
	}
	if Id1 < Id2 {
		for i := Id1 - 1; i <= Id2-1; i++ {
			data2 = append(data2, config.Request{
				Id:        config.RequestData[i].Id,
				Country:   config.RequestData[i].Country,
				Continent: config.RequestData[i].Continent,
				Currency:  config.RequestData[i].Currency,
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"error": false,
			"data":  data2,
		})
		return
	}
}
