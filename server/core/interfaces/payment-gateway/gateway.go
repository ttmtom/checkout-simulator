package payment_gateway

type CoinbaseWebhookEvent struct {
	ID            string    `json:"id"`
	ScheduledFor  string    `json:"scheduled_for"`
	AttemptNumber int       `json:"attempt_number"`
	Event         EventData `json:"event"`
}

type EventData struct {
	ID         string     `json:"id"`
	Resource   string     `json:"resource"`
	Type       string     `json:"type"`
	APIVersion string     `json:"api_version"`
	CreatedAt  string     `json:"created_at"`
	Data       ChargeData `json:"data"`
}

type ChargeData struct {
	Metadata map[string]interface{} `json:"metadata"`
}

type ChargeResponse struct {
	Id         string `json:"id"`
	PaymentUrl string `json:"payment_url"`
}

type Gateway interface {
	CreateCharge(id string) *ChargeResponse
	ValidateEvent(valid bool) bool
}
