package routers

import (
	"github.com/gin-gonic/gin"
	"gin-framework/controllers"
)

func StartServer() *gin.Engine{
	// membuat routers
	routers := gin.Default()

	routers.POST("/addcar", controllers.CreateCar)

	routers.GET("/addcar/:CarID", controllers.GetCar)

	routers.PUT("/addcar/:CarID", controllers.UpdateCar)

	routers.DELETE("/addcar/:CarId", controllers.DeleteCar)

	return routers
}