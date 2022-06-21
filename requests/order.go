package requests

import "assignment/models"

type OrderRequest struct {
	CustomerName string        `json:"customerName" binding:"required"`
	Items        []models.Item `json:"items" binding:"required"`
}
