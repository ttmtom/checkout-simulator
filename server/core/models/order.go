package models

type Order struct {
	ID     int64   `json:"id" db:"id"`
	User   string  `json:"user" db:"user"`
	Amount float64 `json:"amount" db:"amount"`
}
