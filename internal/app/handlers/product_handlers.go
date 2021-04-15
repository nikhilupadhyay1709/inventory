package handlers

import (
	"inventory/configs/database"
	"inventory/internal/app/types"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var product types.Product

	var responseCode int
	responseMessage := ""
	if c.ShouldBind(&product) == nil {
		if errors := database.DBConn.Create(&types.Product{
			Name: product.Name,
			Sku:  product.Sku,
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
