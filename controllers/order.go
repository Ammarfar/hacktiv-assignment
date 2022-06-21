package controllers

import (
	"assignment/configs"
	"assignment/helpers"
	"assignment/models"
	"assignment/requests"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var orderRequest requests.OrderRequest
	if err := c.ShouldBindJSON(&orderRequest); err != nil {
		c.JSON(http.StatusBadRequest, helpers.ResponseError("Bad Request", nil))
		return
	}

	order := models.Order{
		OrderedAt:    time.Now(),
		CustomerName: orderRequest.CustomerName,
		Items:        orderRequest.Items,
	}

	configs.DB.Create(&order)

	c.JSON(http.StatusOK, helpers.ResponseSuccess("Order Created Successfully!", &order))
	return
}

func GetOrder(c *gin.Context) {
	orders := []models.Order{}

	configs.DB.Preload("Items").Find(&orders)

	c.JSON(http.StatusOK, helpers.ResponseSuccess("Success Retrieving Data!", &orders))
	return
}

func UpdateOrder(c *gin.Context) {
	orderRequest := requests.OrderRequest{}
	if err := c.ShouldBindJSON(&orderRequest); err != nil {
		c.JSON(http.StatusBadRequest, helpers.ResponseError("Bad Request", nil))
		return
	}

	orderId := c.Param("orderId")
	order := models.Order{}
	item := models.Item{}

	for _, v := range orderRequest.Items {
		configs.DB.Model(item).Where("id = ?", v.ID).Updates(map[string]interface{}{
			"code":        v.Code,
			"description": v.Description,
			"quantity":    v.Quantity,
		})
	}

	configs.DB.Model(&order).Where("id = ?", &orderId).Updates(models.Order{
		CustomerName: orderRequest.CustomerName,
	})

	c.JSON(http.StatusOK, helpers.ResponseSuccess("Order Updated Successfully!", &order))
	return
}

func DeleteOrder(c *gin.Context) {
	orderId := c.Param("orderId")

	configs.DB.Delete(&models.Order{}, orderId)
	configs.DB.Where("order_id = ?", orderId).Delete(&models.Item{})

	c.JSON(http.StatusOK, helpers.ResponseSuccess("Order Deleted Successfully!", nil))
	return
}
