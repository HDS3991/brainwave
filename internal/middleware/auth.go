package middleware

import (
	"brainwave/pkg/consts/header"
	"github.com/gin-gonic/gin"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get(header.JWTHeaderName)
		if len(token) <= 0 {
			c.Next()
			return
		}
		c.Next()
	}
}

func SessionAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
