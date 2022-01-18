package main

import (
	"github.com/Ledan-bot/groom-room/server/controllers"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	setupServer().Run(":8072")
}

func setupServer() *gin.Engine {
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./public", true)))
	router.GET("/api/test", controllers.HandleTest)
	router.GET("/api/users/all", controllers.GetAllUsers)
	router.POST("/api/users/create", controllers.CreateNewUser)
	router.POST("/api/users/login", controllers.LoginUser)
	router.GET("/api/appointments/all", controllers.GetAllAppointments)
	router.GET("/api/animals/all", controllers.GetAllAnimals)
	return router
}
