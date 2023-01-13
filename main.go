package main

import (
	"crud-api/src/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// here wer run all the routes
	routes.Routes(router)
}
