package api

import (
	"crud-api/config"
	"log"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
)

func ContinentSort(c *gin.Context) {
	var (
		data      []config.Request
		continent []string
	)
	for i := 0; i < len(config.RequestData); i++ {
		continent = append(continent, config.RequestData[i].Continent)
	}
	s := Unique(continent)
	sort.Strings(s)
	log.Println(s)
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(config.RequestData); j++ {
			if s[i] == config.RequestData[j].Continent {
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

func Unique(s []string) []string {
	inResult := make(map[string]bool)
	var result []string
	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	return result
}
