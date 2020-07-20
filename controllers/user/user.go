package user

import (
	"net/http"
	"todo/models"
	"todo/module/session"

	"github.com/gin-gonic/gin"
)

func GetInfoById(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

func Register(c *gin.Context) {
	json := struct {
		Name     string `json:"name" binding:"required,min=2,max=5"`
		Password string `json:"password" binding:"required,password"`
	}{}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			// "error": 1,
			"msg": err.Error(),
		})
		return
	}
	if uid, err := models.CreateUser(json.Name, json.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			// "error": 2,
			"msg": err.Error(),
		})
		return
	} else {
		// 设置 session
		sid := session.NewSession(uid)
		c.SetCookie("sid", sid, 3600, "/", "localhost", false, true)
		c.String(http.StatusOK, "success")
	}

}

func loginByPassword(c *gin.Context) {
	json := struct {
		Name     string `json:"name" binding:"required,min=2,max=5"`
		Password string `json:"password" binding:"required,password"`
	}{}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			// "error": 2,
			"msg": err.Error(),
		})
	}
	if err := models.CheckUsernameAndPassword(json.Name, json.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			// "error": 2,
			"msg": err.Error(),
		})
	}
	// 创建 sessionID，保存到 cookies 里以及 JSON 里。

}
