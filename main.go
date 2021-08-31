package main

import (
	"github.com/camilocorreaUdeA/GoBootcampTechTest/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/hello", handlers.RequestWrapper(handlers.SayHello))

	router.GET("/foo", handlers.RequestWrapper(handlers.GetCustomData))

	router.Run(":8080")
}
