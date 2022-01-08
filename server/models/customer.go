package models

type Customer struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	AnimalID  int64  `json:"animal_id"`
}
