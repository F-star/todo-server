package todo

import (
	"net/http"
	"strconv"
	"todo/dto"
	"todo/models"

	"github.com/gin-gonic/gin"
)

// const TodoCollection = "todo"

func CreateATodo(c *gin.Context) {
	type Json struct {
		Title string `json:"title" binding:"required"`
	}

	var json Json
	// 参数校验
	// 如果 JSON 为空字符串，err.Error() 会变成 EOF
	if paramsErr := c.ShouldBindJSON(&json); paramsErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": 1,
			"msg":   paramsErr.Error(),
		})
		return
	}

	// 数据库操作
	// title = json.Title
	_, err := models.CreateATodo(json.Title)
	if err != nil {
		// c.AbortWithStatus(http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": 2,
			"msg":   err.Error(),
		})
	} else {
		c.Status(http.StatusOK)
	}
}

func GetATodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo

	if err := models.GetATodo(&todo, id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": 2,
			"msg":   err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, dto.ToTodoDTO(todo))
	}
}

func GetAllTodos(c *gin.Context) {
	todos, err := models.GetAllTodos()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": 2,
			"msg":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, dto.ToTodoDTOs(todos))
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	var rowsAffected int64
	var err error
	if rowsAffected, err = models.DeleteATodo(id); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	c.String(http.StatusOK, "updated num: "+strconv.FormatInt(rowsAffected, 10))
}

func UpdateATodo(c *gin.Context) {
	type Json struct {
		ID    uint   `json:"id" binding:"required"`
		Title string `json:"title" binding:"required"`
	}
	var json Json
	if err := c.ShouldBind(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": 2,
			"msg": err.Error(),
		})
		return
	}

	if err := models.UpdateATodo(json.ID, json.Title); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": 2,
			"msg": err.Error(),
		})
		return
	}
	c.Status(http.StatusOK)
}
