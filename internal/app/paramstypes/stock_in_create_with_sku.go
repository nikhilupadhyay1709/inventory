package paramstypes

type StockInCreateWithSku struct {
	PricePerProduct   float32 `json:"price_per_product"`
	TransactionNumber string  `json:"transaction_number"`
	Note              string  `json:"note"`
	Sku               string  `json:"sku"`
	OrderedQuantity   int     `json:"ordered_quantity"`
	ReceivedQuantity  int     `json:"received_quantity"`
}
