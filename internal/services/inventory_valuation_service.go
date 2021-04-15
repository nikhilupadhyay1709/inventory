package services

import (
	"inventory/configs/database"
	"inventory/internal/app/reporttypes"
	"inventory/internal/app/types"
)

// A service to calculate inventory valuation.
// Find all product.
// For each product get stock ins quantities and total prices.
func InventoryValuationCalculate() []reporttypes.InventoryValuation {
	db := database.DBConn
	var products []types.Product
	var inventoryValuations []reporttypes.InventoryValuation

	db.Preload("StockIns").Find(&products)

	for _, product := range products {
		var inventoryValuation = reporttypes.InventoryValuation{}
		inventoryValuation.ProductSku = product.Sku
		inventoryValuation.ProductName = product.Name

		var grandTotalPrice float32

		totalQuantityStockIn := 0
		grandTotalPrice = 0.0
		for _, stockIn := range product.StockIns {
			totalQuantityStockIn += stockIn.ReceivedQuantity
			grandTotalPrice += float32(stockIn.ReceivedQuantity) * stockIn.PricePerProduct
		}

		inventoryValuation.ProductQuantity = totalQuantityStockIn
		inventoryValuation.ProductTotalPurchasePrice = grandTotalPrice
		inventoryValuation.ProductAvgPurchasePrice = divideByPossibleZero(grandTotalPrice,
			totalQuantityStockIn)

		inventoryValuations = append(inventoryValuations, inventoryValuation)
	}

	return inventoryValuations
}
