# Backend Engineer Case Study: Crypto Checkout Simulator

**Background:**
Build a lightweight simulation of a crypto checkout backend.

### 1. API Endpoint
- **POST** `/checkout`
- Accepts: `{ "amount": `number`, "email": `string` }`
- Simulates calling the Coinbase Commerce API (mock it) and returns a fake hosted payment url (e.g. `https://fake.coinbase.com/pay/12345` )

### 2. Webhook Receiver
- **POST** `/webhook`
- Accepts a simulated Coinbase webhook (mock payload) and:
    - Validates the payload
    - Logs or stores the transaction (status, email, timestamp, fake transaction ID)

