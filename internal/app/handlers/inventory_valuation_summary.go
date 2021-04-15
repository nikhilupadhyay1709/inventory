package handlers

import (
	"fmt"
	"inventory/configs/database"
	"inventory/internal/app/reporttypes"
	"inventory/internal/app/types"
	"inventory/internal/services"

	"github.com/gin-gonic/gin"
)

func InventoryValuationSummary(c *gin.Context) {
	db := database.DBConn
	var inventoryValuationSummary reporttypes.InventoryValuationSummary

	var productCount int
	db.Model(&types.Product{}).Count(&productCount)

	var totalStockIn reporttypes.SumResult
	db.Model(&types.StockIn{}).Select("sum(received_quantity) as sum").Scan(&totalStockIn)

	fmt.Println(totalStockIn)

	var totalStockOut reporttypes.SumResult
	db.Model(&types.StockOut{}).Select("sum(quantity) as sum").Scan(&totalStockOut)

	fmt.Println(totalStockOut)
	productTotalQuantity := totalStockIn.Sum - totalStockOut.Sum

	inventoryValuations := services.InventoryValuationCalculate()
	var totalValuation float32 = 0.0
	for _, iv := range inventoryValuations {
		totalValuation += iv.ProductTotalPurchasePrice
	}

	inventoryValuationSummary.ProductSkuCount = productCount
	inventoryValuationSummary.ProductTotalQuantity = productTotalQuantity
	inventoryValuationSummary.TotalValuation = totalValuation

	//json, err := json.Marshal(inventoryValuationSummary)

	c.JSON(200, inventoryValuationSummary)
}
