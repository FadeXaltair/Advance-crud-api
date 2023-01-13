package api

import (
	"crud-api/config"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var body config.Request
	err := c.Bind(&body)
	if err != nil {
		log.Println(err)
		return
	}
	for i := 0; i < len(config.RequestData); i++ {
		if config.RequestData[i].Id == body.Id {
			c.JSON(http.StatusOK, gin.H{
				"error":   true,
				"message": "Id cannot be same",
			})
			return
		}
	}
	config.RequestData = append(config.RequestData, config.Request{
		Id:        body.Id,
		Country:   body.Country,
		Continent: body.Continent,
		Currency:  body.Currency,
	})

	data := config.Request{
		Id:        body.Id,
		Country:   body.Country,
		Continent: body.Continent,
		Currency:  body.Currency,
	}
	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  data,
	})
}
