package handlers

import (
	"inventory/configs/database"
	"inventory/internal/app/paramstypes"
	"inventory/internal/app/types"
	"time"

	"github.com/gin-gonic/gin"
)

// First check if received quantity > ordered quantity.
// If yes, return 422.
//
// Second find if product by such SKU exists.
// If not return response 422.
//
// If product already exists, add the quantity.
// And create stock in data.
func CreateStockIn(c *gin.Context) {
	db := database.DBConn
	var responseCode int
	responseMessage := ""
	var StockInParams paramstypes.StockInCreateWithSku

	if c.ShouldBind(&StockInParams) == nil {
		// =============================================================================
		// VALIDATIONS
		// =============================================================================
		if StockInParams.ReceivedQuantity > StockInParams.OrderedQuantity {
			c.String(422, "Received Quantity cannot be more than Ordered Quantity")
			return
		}

		var stockInWithSimilarTransactionNumber types.StockIn
		if err := db.Where("transaction_number = ?", StockInParams.TransactionNumber).
			Find(&stockInWithSimilarTransactionNumber).Error; err == nil {
			c.String(422, "Transaction Number already exists")
			return
		}

		var product types.Product
		if err := db.Where("sku = ?", StockInParams.Sku).First(&product).Error; err != nil {
			c.String(422, "Product does not exist")
			return
		}
		// =============================================================================

		if errors := db.Create(&types.StockIn{
			PricePerProduct:   StockInParams.PricePerProduct,
			TransactionNumber: StockInParams.TransactionNumber,
			OrderedQuantity:   StockInParams.OrderedQuantity,
			ReceivedQuantity:  StockInParams.ReceivedQuantity,
			Product:           product,
			Note:              StockInParams.Note,
			Time:              time.Now(),
		}).GetErrors(); len(errors) > 0 {
			responseCode = 422
			for _, err := range errors {
				responseMessage = responseMessage + ", " + err.Error()
			}
		} else {
			responseCode = 201
			responseMessage = "Created"
		}
	}

	c.String(responseCode, responseMessage)
}
