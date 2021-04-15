package handlers

import (
	"inventory/configs/database"
	"inventory/internal/app/types"

	"github.com/gin-gonic/gin"
)

func DeleteStockOut(c *gin.Context) {
	db := database.DBConn
	stockOut := types.StockOut{}
	id := c.Param("id")

	// =============================================================================
	// VALIDATIONS
	// =============================================================================
	if err := db.First(&stockOut, id).Error; err != nil {
		c.String(404, "Stock Out Not Found")
		return
	}

	db.Unscoped().Delete(&stockOut)

	c.String(200, "Stock Out Deleted")
}
