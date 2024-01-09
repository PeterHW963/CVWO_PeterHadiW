package main

import (
	"net/http"

	"github.com/PeterHW963/CVWO/backend/config"
	"github.com/PeterHW963/CVWO/backend/route"
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

func main() {
	router := gin.New()

	opts := cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	}

	handler := cors.New(opts).Handler(router)

	config.Connect()
	route.UserRoute(router)

	http.ListenAndServe(":8080", handler)
}
