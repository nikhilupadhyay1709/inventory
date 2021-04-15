package services

import (
	"inventory/configs/database"
	"inventory/internal/app/reporttypes"
	"inventory/internal/app/types"
	"time"
)

// Get inventory valuations in order to get buy price.
// Then for each stock out, obtain product buy price, then calculate the profit/loss.
func SalesReportCalculate(startDateString string, endDateString string) []reporttypes.SalesReport {
	db := database.DBConn
	inventoryValuations := InventoryValuationCalculate()
	var salesReports []reporttypes.SalesReport
	var stockOuts []types.StockOut

	skuAndAvgBuyPrices := make(map[string]float32)
	for _, iv := range inventoryValuations {
		skuAndAvgBuyPrices[iv.ProductSku] = iv.ProductAvgPurchasePrice
	}

	startDate, _ := time.Parse(time.RFC3339, startDateString+"T00:00:00.000Z")
	endDate, _ := time.Parse(time.RFC3339, endDateString+"T23:59:59.000Z")

	db.Preload("Product").
		Where("created_at BETWEEN ? AND ?", startDate, endDate).
		Find(&stockOuts)

	for _, stockOut := range stockOuts {
		var salesReport = reporttypes.SalesReport{}
		salesReport.SalesNote = stockOut.Note
		salesReport.Time = stockOut.Time
		salesReport.ProductSku = stockOut.Product.Sku
		salesReport.ProductName = stockOut.Product.Name
		salesReport.Quantity = stockOut.Quantity

		// Sales prices
		salesReport.SellPricePerProduct = stockOut.PricePerProduct
		totalSellPrice := stockOut.PricePerProduct * float32(stockOut.Quantity)
		salesReport.TotalSellPricePerProduct = totalSellPrice

		// Buy prices
		buyPricePerProduct := skuAndAvgBuyPrices[stockOut.Product.Sku]
		salesReport.BuyPricePerProduct = buyPricePerProduct
		totalBuyPrice := buyPricePerProduct * float32(stockOut.Quantity)
		salesReport.TotalBuyPricePerProduct = totalBuyPrice

		salesReport.ProfitOrLoss = totalSellPrice - totalBuyPrice
		salesReports = append(salesReports, salesReport)
	}

	return salesReports
}
