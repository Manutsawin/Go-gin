package main

import (
	"github.com/gin-gonic/gin"
	"github.com/itsmaheshkariya/gin-gorm-rest/config"
	"github.com/itsmaheshkariya/gin-gorm-rest/routes"
)

func main() {
	router := gin.New()
	config.Connect()
	routes.UserRoute(router)
	router.Run(":8080")
}
