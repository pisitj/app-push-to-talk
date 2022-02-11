package middleware

import (
	"github.com/gin-gonic/gin"
)

func CheckUserLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

func CheckUserWhitelist() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

func CheckUserAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
