package login

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func abortWithError(c *gin.Context, code int, message string) {
	c.AbortWithStatusJSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}

func NewRoutesFactory(group *gin.RouterGroup) func(service Service) {
	loginFactory := func(loginService Service) {
		group.POST("/", func(c *gin.Context) {
			var user User

			if err := c.BindJSON(&user); err != nil {
				abortWithError(c, http.StatusBadRequest, "Missing user field.")
			}

			if !loginService.IsValidUser(user) {
				abortWithError(c, http.StatusUnauthorized, "User is unauthorized.")
			}

			c.Status(200)
		})
	}
	return loginFactory
}
