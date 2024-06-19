package models

type Order struct {
    ID         uint    `json:"id"`
    UserID     uint    `json:"user_id"`
    ProductID  uint    `json:"product_id"`
    Quantity   int     `json:"quantity"`
    TotalPrice float64 `json:"total_price"`
    Status     string  `json:"status"`
}
