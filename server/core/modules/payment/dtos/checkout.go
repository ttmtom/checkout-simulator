package dtos

type CheckoutDto struct {
	Email  string  `json:"email" validate:"required,email"`
	Amount float64 `json:"amount" validate:"required,gt=0"`
}
