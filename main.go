package main

import (
	"github.com/Ledan-bot/groom-room/server/controllers"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	setupServer().Run(":8070")
}

func setupServer() *gin.Engine {
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./public", true)))
	router.GET("/api/test", controllers.HandleTest)
	router.GET("/api/appointments/all", controllers.GetAllAppointments)
	router.GET("/api/customers/all", controllers.GetAllCustomers)
	router.GET("/api/animals/all", controllers.GetAllAnimals)
	router.POST("/api/customer/new", controllers.CreateNewCustomer)
	router.POST("/api/login", controllers.LoginUser)
	return router
}
