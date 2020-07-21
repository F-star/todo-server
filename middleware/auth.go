package routehandler

import (
	"log"
	"net/http"
	"todo/module/session"

	"github.com/gin-gonic/gin"
)

func Authhandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("auth 中间件")
		// TODO: check token
		sid, err := c.Cookie("sid")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{})
			return
		}
		uid := session.GetUIDBySID(sid)
		if uid == "" {
			c.JSON(http.StatusUnauthorized, gin.H{}) // expired or wrong sid
			return
		}
		c.Set("uid", uid)
		c.Next()
	}
}
