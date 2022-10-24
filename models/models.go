package models

type Stock struct {
	StockID int64   `json:"stockID"`
	Name    string  `json:"name"`
	Price   float64 `json:"price"`
	Company string  `json:"company"`
}
