package models

type Appointment struct {
	ID         int64  `json:"id"`
	CustomerID int64  `json:"customer_id"`
	Date       string `json:"date"`
	AnimalID   int64  `json:"animal_id"`
}
