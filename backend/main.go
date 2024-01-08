package main

import (
	"github.com/PeterHW963/CVWO/backend/config"
	"github.com/PeterHW963/CVWO/backend/route"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	config.Connect()
	route.UserRoute(router)

	router.Run(":8080")
}
