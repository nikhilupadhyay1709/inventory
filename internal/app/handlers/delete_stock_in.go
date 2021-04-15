package handlers

import (
	"inventory/configs/database"
	"inventory/internal/app/types"

	"github.com/gin-gonic/gin"
)

func DeleteStockIn(c *gin.Context) {
	db := database.DBConn
	stockIn := types.StockIn{}
	id := c.Param("id")

	// =============================================================================
	// VALIDATIONS
	// =============================================================================
	if err := db.First(&stockIn, id).Error; err != nil {
		c.String(404, "Stock In Not Found")
		return
	}

	db.Unscoped().Delete(&stockIn)

	c.String(200, "Stock In Deleted")
}
