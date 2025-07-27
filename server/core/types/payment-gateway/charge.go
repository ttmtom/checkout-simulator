package payment_gateway

type ChargeResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Amount string `json:"amount"`
}
