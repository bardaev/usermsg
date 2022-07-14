package main

import (
	"usermsg/controller"
	"usermsg/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	const conn string = "user=go password=go dbname=go sslmode=disable port=5430"

	repo := repository.NewUserRepo(conn)
	controller := controller.NewUserController(repo)

	router := gin.Default()

	router.POST("/user", controller.CreateUser)
	// router.GET("/user/:name", controller.)

	router.Run()
}
