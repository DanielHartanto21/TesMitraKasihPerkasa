package web

type EditStockRequest struct {
	ProductID int `json:"productID"`
	Stock     int `json:"stock"`
}
