package routes

import (
	"crud-api/src/api"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.POST("/create", api.Create)
	router.GET("/read", api.Read)
	router.PUT("/update", api.Update)
	router.DELETE("/delete", api.Delete)
	router.GET("/id-sort", api.Sort)
	router.GET("/filter", api.Filter)
	router.GET("/country-sort", api.CountrySort)
	router.GET("/details", api.Details)
	router.Run(":3000")
}
