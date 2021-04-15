package handlers

import (
	"inventory/configs/database"
	"inventory/internal/app/types"

	"github.com/gin-gonic/gin"
)

// Cannot delete product if already have stock-in or stock-out.
func DeleteProduct(c *gin.Context) {
	db := database.DBConn
	product := types.Product{}
	id := c.Param("id")

	// =============================================================================
	// VALIDATIONS
	// =============================================================================
	if err := db.First(&product, id).Error; err != nil {
		c.String(404, "Product Not Found")
		return
	}

	productId := product.ID
	// Check stock-in.
	stockIn := types.StockIn{}
	var hasStockIn bool = false
	if err := db.Where("product_id = ?", productId).First(&stockIn).Error; err == nil {
		hasStockIn = true
	}

	// Check stock-out
	stockOut := types.StockOut{}
	var hasStockOut bool = false
	if err := db.Where("product_id = ?", productId).First(&stockOut).Error; err == nil {
		hasStockOut = true
	}

	if hasStockIn && hasStockOut {
		c.String(422, "Cannot delete product, already has stock-in and out")
		return
	} else if hasStockIn {
		c.String(422, "Cannot delete product, already has stock-in")
		return
	} else if hasStockOut {
		c.String(422, "Cannot delete product, already has stock-out")
		return
	}

	// =============================================================================

	db.Unscoped().Delete(&product)

	c.String(200, "Product Deleted")
}
