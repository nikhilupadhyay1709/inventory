package handlers

import (
	"inventory/configs/database"
	"inventory/internal/app/paramstypes"
	"inventory/internal/app/types"

	"github.com/gin-gonic/gin"
)

func UpdateProduct(c *gin.Context) {
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

	// =============================================================================

	// For now user can only update product name.
	productParams := paramstypes.UpdateProductParam{}
	if c.ShouldBind(&productParams) == nil {
		product.Name = productParams.Name
		db.Save(&product)
	}

	c.String(200, "Product Updated")
}
