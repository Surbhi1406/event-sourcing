package models

type Order struct {
	OrderID  string `json:"orderId"`
	Item     string `json:"item"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
}
