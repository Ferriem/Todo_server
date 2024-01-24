package controllers

import (
	"github.com/Ferriem/Todo_server/models"
	"github.com/Ferriem/Todo_server/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type UserController struct{}

func (uc *UserController) Register(c *gin.Context) {
	var auth models.Auth
	if err := c.ShouldBindJSON(&auth); err != nil {
		c.JSON(500, gin.H{"error": "Invalid json"})
		return
	}
	if auth.Password == "" {
		c.JSON(500, gin.H{
			"message": "Please input your password",
		})
		return
	}
	err := utils.NewUser(auth.ID, auth.Password)
	switch err {
	case utils.ErrUserExists:
		c.JSON(500, gin.H{
			"message": "User already exists, please try another username",
		})
	case utils.ErrServer:
		c.JSON(500, gin.H{
			"message": "Server error",
		})
	default:
		c.JSON(200, gin.H{
			"message": "Success Register",
		})
	}
}

func (uc *UserController) GetFirst(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindBodyWith(&todo, binding.JSON); err != nil {
		c.JSON(500, gin.H{"error": "Invalid json"})
		return
	}
	value, err := utils.GetFirst(todo.ID)
	if err != utils.NoError {
		c.JSON(500, gin.H{
			"message": "Server Error",
		})
		return
	}
	c.JSON(200, gin.H{
		"FirstTodo": value,
	})
}

func (uc *UserController) GetInfo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindBodyWith(&todo, binding.JSON); err != nil {
		c.JSON(500, gin.H{"error": "Invalid json"})
		return
	}
	value, err := utils.GetInfo(todo.ID)
	if err != utils.NoError {
		c.JSON(500, gin.H{
			"message": "Server Error",
		})
		return
	}
	if len(value) == 0 {
		c.JSON(200, gin.H{
			"message": "Empty",
		})
		return
	}
	c.JSON(200, gin.H{
		"Todo": value,
	})
}

func (uc *UserController) Add(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindBodyWith(&todo, binding.JSON); err != nil {
		c.JSON(500, gin.H{"error": "Invalid json"})
		return
	}
	err := utils.Add(todo.ID, todo.Title, todo.Description)
	if err != utils.NoError {
		c.JSON(500, gin.H{
			"message": "Error Add",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Success Add",
	})
}

func (uc *UserController) Delete(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindBodyWith(&todo, binding.JSON); err != nil {
		c.JSON(500, gin.H{"error": "Invalid json"})
		return
	}
	err := utils.Delete(todo.ID)
	if err != utils.NoError {
		c.JSON(500, gin.H{
			"message": "Error Delete",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Success Delete",
	})
}

func (uc *UserController) Update(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindBodyWith(&todo, binding.JSON); err != nil {
		c.JSON(500, gin.H{"error": "Invalid json"})
		return
	}
	err := utils.Update(todo.ID, todo.Title, todo.Description)
	if err != utils.NoError {
		c.JSON(500, gin.H{
			"message": "Error Update",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Success Update",
	})
}
