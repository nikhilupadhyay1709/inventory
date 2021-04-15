package handlers

import (
	"inventory/configs/database"
	"inventory/internal/app/paramstypes"
	"inventory/internal/app/types"
	"time"

	"github.com/gin-gonic/gin"
)

func UpdateStockOut(c *gin.Context) {
	db := database.DBConn
	var responseCode int
	responseMessage := ""

	var stockOutParams paramstypes.StockOutWithParams

	if c.ShouldBind(&stockOutParams) == nil {
		// =============================================================================
		// VALIDATIONS
		// =============================================================================
		var stockOut = types.StockOut{}

		id := c.Param("id")
		if err := db.First(&stockOut, id).Error; err != nil {
			c.String(404, "Stock Out Not Found")
			return
		}

		var product types.Product
		if err := db.Where("sku = ?", stockOutParams.Sku).First(&product).Error; err != nil {
			c.String(422, "Product does not exist")
			return
		}

		// =============================================================================

		if errors := db.Model(&stockOut).Updates(&types.StockOut{
			PricePerProduct: stockOutParams.PricePerProduct,
			Quantity:        stockOutParams.Quantity,
			Product:         product,
			Note:            stockOutParams.Note,
			Time:            time.Now(),
		}).GetErrors(); len(errors) > 0 {
			responseCode = 422
			for _, err := range errors {
				responseMessage = responseMessage + ", " + err.Error()
			}
		} else {
			responseCode = 201
			responseMessage = "Updated"
		}
	}

	c.String(responseCode, responseMessage)
}
