# Backend Engineer Case Study: Crypto Checkout Simulator

**Background:**
You're taking over a backend service that powers a crypto checkout portal for event tickets. Currently, it uses PHP and Coinbase Commerce to create payment sessions and track incoming transactions. Transaction records are stored in a relational database (likely MySQL) and the service is hosted on DigitalOcean.
The existing codebase is hard to extend, poorly documented, and you've been brought in to stabilise and modernise it over time.

## Your Task (2-4 hours expected):
Build a lightweight simulation of a crypto checkout backend using any stack you're comfortable with (PHP preferred, but not required). Your submission should include:

### 1. API Endpoint
- **POST** `/checkout`
- Accepts: `{ "amount": `number`, "email": `string` }`
- Simulates calling the Coinbase Commerce API (mock it) and returns a fake hosted payment url (e.g. `https://fake.coinbase.com/pay/12345` )

### 2. Webhook Receiver
- **POST** `/webhook`
- Accepts a simulated Coinbase webhook (mock payload) and:
    - Validates the payload
    - Logs or stores the transaction (status, email, timestamp, fake transaction ID)

### 3. Database Design
- Provide a simple schema or SQL dump showing how you'd structure the `transactions` table.

### 4. README
- Briefly explain:
    - How to run your project
    - Assumptions made
    - How you'd improve this if given more time (e.g. error handling, signature verification, polling, retries, dashboards, etc.)

## Bonus (optional, not required):
- Use Docker or make it deployable on DigitalOcean/Vercel/etc.
- Include retry logic or error tracking in the webhook flow
- Provide logs or metrics (e.g. simple console logs or a `/health` endpoint)
- Use proper async handling for webhooks (background processing, job queues)

## What We're Looking For:
- Code clarity and structure
- Understanding of payment flows and webhook-based integrations
- Sound data modeling
- Basic security and validation thinking
- Documentation and communication skills
