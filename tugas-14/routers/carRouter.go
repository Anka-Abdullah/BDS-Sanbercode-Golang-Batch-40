package routers

import (
	"BDS-Sanbercode-Golang-Batch-40/tugas-14/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/cars", controllers.CreateCar)
	router.PUT("cars/:carID", controllers.Updatecar)
	router.GET("/cars/:carID", controllers.GetCar)
	router.DELETE("/cars/:carID", controllers.DeleteCar)

	return router
}
