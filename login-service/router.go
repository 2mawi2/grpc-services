package main

import (
	"github.com/gin-gonic/gin"
	"github.com/marius/grpc-services/login"
)

func SetupRouter(
	loginService login.Service,
) *gin.Engine {
	router := gin.Default()

	loginGroup := router.Group("/login")
	login.NewRoutesFactory(loginGroup)(loginService)

	return router
}
