package models

type PaymentEvent struct {
	PaymentID string  `json:"payment_id"`
	OrderID   uint    `json:"order_id"`
	UserID    uint    `json:"user_id"`
	Email     string  `json:"email"`
	Amount    float64 `json:"amount"`
	Date      string  `json:"date"`
}

// type PaymentEvent struct {
// 	PaymentID string  `bson:"payment_id"`
// 	OrderID   uint    `bson:"order_id"`
// 	UserID    uint    `bson:"user_id"`
// 	Email     string  `bson:"email"`
// 	Amount    float64 `bson:"amount"`
// 	Date      string  `bson:"date"`
// }
