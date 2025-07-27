package payment_gateway

type ChargeResponse struct {
	Id         string `json:"id"`
	PaymentUrl string `json:"payment_url"`
}

type Gateway interface {
	CreateCharge(id string) *ChargeResponse
}
