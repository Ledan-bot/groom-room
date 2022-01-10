package models

type Animal struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Breed      string `json:"breed"`
	CustomerID int64  `json:"custumer_id"`
}
