CREATE TYPE transaction_status AS ENUM (
    'NEW',
    'PENDING',
    'COMPLETED',
    'EXPIRED',
    'UNRESOLVED',
    'CANCELED',
    'FAILED'
    );

CREATE TABLE transactions (
    id BIGSERIAL PRIMARY KEY,
    provider_type TEXT NOT NULL,
    service_provider_id TEXT NOT NULL UNIQUE,
    customer_email TEXT,
    status transaction_status NOT NULL,
    amount NUMERIC(19, 8),
    currency VARCHAR(8),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    last_event_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_transactions_on_service_provider_id ON transactions(service_provider_id);


CREATE TABLE transaction_events (
    id BIGSERIAL PRIMARY KEY,

    provider_type TEXT NOT NULL,
    transaction_id BIGINT NOT NULL REFERENCES transactions(id) ON DELETE CASCADE,
    service_provider_id TEXT NOT NULL UNIQUE,
    status TEXT NOT NULL,
    event_payload JSONB,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_transaction_events_on_transaction_id ON transaction_events(transaction_id);
CREATE INDEX idx_transaction_events_on_service_provider_id ON transaction_events(service_provider_id);
