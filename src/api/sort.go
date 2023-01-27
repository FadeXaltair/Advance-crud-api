package api

import (
	"crud-api/config"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
)

func Namesort(c *gin.Context) {
	key, _ := c.GetQuery("key")
	order, _ := c.GetQuery("dir")
	if key == "" || order == "" {
		c.JSON(400, gin.H{
			"error":   true,
			"message": "mismatch or empty query",
		})
		return
	}

	if order == "asc" {
		if key == "country" {
			sort.Slice(config.RequestData, func(i, j int) bool {
				return config.RequestData[i].Country < config.RequestData[j].Country
			})
		} else if key == "continent" {
			sort.Slice(config.RequestData, func(i, j int) bool {
				return config.RequestData[i].Continent < config.RequestData[j].Continent
			})
		} else if key == "currency" {
			sort.Slice(config.RequestData, func(i, j int) bool {
				return config.RequestData[i].Currency < config.RequestData[j].Currency
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"error": false,
			"data":  config.RequestData,
		})
		return
	} else if order == "desc" {
		if key == "country" {
			sort.Slice(config.RequestData, func(i, j int) bool {
				return config.RequestData[i].Country > config.RequestData[j].Country
			})
		} else if key == "continent" {
			sort.Slice(config.RequestData, func(i, j int) bool {
				return config.RequestData[i].Continent > config.RequestData[j].Continent
			})
		} else if key == "currency" {
			sort.Slice(config.RequestData, func(i, j int) bool {
				return config.RequestData[i].Currency > config.RequestData[j].Currency
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"error": false,
			"data":  config.RequestData,
		})
		return
	}
}
