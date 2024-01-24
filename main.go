package main

import (
	"github.com/Ferriem/Todo_server/config"
	"github.com/Ferriem/Todo_server/controllers"
	"github.com/Ferriem/Todo_server/middleware"
	"github.com/Ferriem/Todo_server/utils"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func main() {
	config.Rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	router := gin.Default()

	userController := &controllers.UserController{}

	Group := router.Group("/user")
	router.Use(utils.LoggerToRedis())
	Group.Use(middleware.AuthMiddleware())

	Group.Any("/info", func(c *gin.Context) {
		userController.GetInfo(c)
	})

	Group.Any("/first", func(c *gin.Context) {
		userController.GetFirst(c)
	})
	router.Any("/register", func(c *gin.Context) {
		userController.Register(c)
	})
	Group.Any("/add", func(c *gin.Context) {
		userController.Add(c)
	})
	Group.Any("/delete", func(c *gin.Context) {
		userController.Delete(c)
	})
	Group.Any("/update", func(c *gin.Context) {
		userController.Update(c)
	})
	router.Run(":8080")
}
