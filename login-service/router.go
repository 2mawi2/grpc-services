package main

import (
	"github.com/gin-gonic/gin"
	"github.com/marius/moduleProject/login"
)

func setupRouter(
	loginService login.Service,
) *gin.Engine {
	router := gin.Default()

	loginGroup := router.Group("/login")
	login.NewRoutesFactory(loginGroup)(loginService)

	return router
}
