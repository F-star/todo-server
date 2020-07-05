package user

import (
	"net/http"
	"todo/models"

	"github.com/gin-gonic/gin"
)

func GetInfoById(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

func CreateUser(c *gin.Context) {
	json := struct {
		Name     string `json:"name" binding:"required,min=2,max=5"`
		Password string `json:"password" binding:"required,password"`
	}{}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": 1,
			"msg":   err.Error(),
		})
		return
	}
	if err := models.CreateUser(json.Name, json.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": 2,
			"msg":   err.Error(),
		})
		return
	}
	c.String(http.StatusOK, "success")
}
