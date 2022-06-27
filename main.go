package main

import (
	"assignment/configs"
	"assignment/controllers"
	"assignment/helpers"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {

	// Assignment 1
	Biodata()

	r := gin.Default()

	configs.ConnectDB()
	sqlDB, err := configs.DB.DB()
	helpers.PanicIfError(err, "error retrieving DB func")
	defer sqlDB.Close()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, helpers.ResponseSuccess("pong", nil))
	})

	// Assignment 2
	r.POST("/orders", controllers.CreateOrder)
	r.GET("/orders", controllers.GetOrder)
	r.PUT("/orders/:orderId", controllers.UpdateOrder)
	r.DELETE("/orders/:orderId", controllers.DeleteOrder)

	host := helpers.GetEnv("APP_HOST")
	port := helpers.GetEnv("APP_PORT")
	r.Run(host + ":" + port)
}

func Biodata() {
	input, _ := strconv.Atoi(os.Args[1])
	user := helpers.GenerateUser(5)

	if input > len(user) {
		panic("User Not Found")
	}

	fmt.Printf("%+v\n", user[input-1])
}
