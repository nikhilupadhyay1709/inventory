package routes

import (
	"inventory/internal/app/handlers"

	"github.com/gin-gonic/gin"
)

func GenerateRoutes() *gin.Engine {
	web := gin.Default()
	web.GET("/ping", handlers.HandlePing)

	web.GET("v1/products", handlers.ProductList)
	web.POST("v1/products", handlers.CreateProduct)
	web.PATCH("v1/products/:id", handlers.UpdateProduct)
	web.DELETE("v1/products/:id", handlers.DeleteProduct)

	web.POST("v1/stock_ins", handlers.CreateStockIn)
	web.PATCH("v1/stock_ins/:id", handlers.UpdateStockIn)
	web.DELETE("v1/stock_ins/:id", handlers.DeleteStockIn)

	web.POST("v1/stock_outs", handlers.CreateStockOut)
	web.PATCH("v1/stock_outs/:id", handlers.UpdateStockOut)
	web.DELETE("v1/stock_outs/:id", handlers.DeleteStockOut)

	web.GET("v1/reports/inventory_valuation_summary", handlers.InventoryValuationSummary)
	web.GET("v1/reports/sales_report_summary", handlers.SalesReportSummary)

	web.GET("v1/reports/inventory_valuation.csv", handlers.InventoryValuation)
	web.GET("v1/reports/sales_report.csv", handlers.SalesReport)

	return web
}
