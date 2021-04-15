package handlers

import (
	"inventory/configs/database"
	"inventory/internal/app/types"

	"github.com/gin-gonic/gin"
)

func ProductList(c *gin.Context) {
	db := database.DBConn
	var products []types.Product

	db.Preload("StockIns").Preload("StockOuts").
		Find(&products)

	var productsList []types.ProductInView

	for _, product := range products {
		currentQuantity := calculateCurrentQuantity(product)
		productShow := types.ProductInView{
			Sku:             product.Sku,
			Name:            product.Name,
			CurrentQuantity: currentQuantity,
		}

		productsList = append(productsList, productShow)
	}

	c.JSON(200, productsList)
}

func calculateCurrentQuantity(product types.Product) int {
	totalStockInsQuantity := 0
	totalStockOutsQuantity := 0

	for _, stockIn := range product.StockIns {
		totalStockInsQuantity += stockIn.ReceivedQuantity
	}

	for _, stockOut := range product.StockOuts {
		totalStockOutsQuantity += stockOut.Quantity
	}

	return totalStockInsQuantity - totalStockOutsQuantity
}
