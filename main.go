package main

import (
	"assignment/configs"
	"assignment/controllers"
	"assignment/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	configs.ConnectDB()
	sqlDB, err := configs.DB.DB()
	helpers.PanicIfError(err, "error retrieving DB func")
	defer sqlDB.Close()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, helpers.ResponseSuccess("pong", nil))
	})

	r.POST("/orders", controllers.CreateOrder)
	r.GET("/orders", controllers.GetOrder)
	r.PUT("/orders/:orderId", controllers.UpdateOrder)
	r.DELETE("/orders/:orderId", controllers.DeleteOrder)

	host := helpers.GetEnv("APP_HOST")
	port := helpers.GetEnv("APP_PORT")
	r.Run(host + ":" + port)
}
