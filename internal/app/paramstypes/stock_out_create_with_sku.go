package paramstypes

type StockOutWithParams struct {
	Sku             string  `json:"sku"`
	PricePerProduct float32 `json:"price_per_product"`
	Note            string  `json:"note"`
	Quantity        int     `json:"quantity"`
}
