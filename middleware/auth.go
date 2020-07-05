package routehandler

import (
	"log"

	"github.com/gin-gonic/gin"
)

func AuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("auth 中间件")
		// TODO: check token
		c.Next()
	}
}
