BEGIN;
    CREATE TYPE payment_status_enum AS ENUM (
        'NEW',
        'CREATED',
        'PENDING',
        'COMPLETED',
        'FAILED'
    );

    CREATE TYPE order_status_enum AS ENUM (
        'PENDING',
        'PROCESSING',
        'COMPLETED',
        'FAILED'
    );

    CREATE TABLE orders (
        id BIGSERIAL PRIMARY KEY,
        "user" VARCHAR(255) NOT NULL,
        amount NUMERIC(19, 4) NOT NULL,
        status order_status_enum NOT NULL
    );

    CREATE TABLE payments (
        id BIGSERIAL PRIMARY KEY,
        order_id BIGINT NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
        service_provider_id VARCHAR(255) NOT NULL,
        payment_url VARCHAR(255) NOT NULL,
        status payment_status_enum NOT NULL,
        amount NUMERIC(19, 4) NOT NULL,
        created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
        last_event_at TIMESTAMPTZ
    );

    CREATE INDEX idx_payments_order_id ON payments(order_id);
    CREATE INDEX idx_payments_service_provider_id ON payments(service_provider_id);
    CREATE INDEX idx_payments_status ON payments(status);

    CREATE TABLE payment_events (
        id BIGSERIAL PRIMARY KEY,
        payment_id BIGINT NOT NULL REFERENCES payments(id) ON DELETE CASCADE,
        service_provider_id VARCHAR(255) NOT NULL,
        status payment_status_enum NOT NULL,
        event_payload JSONB,
        created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
    );

    CREATE INDEX idx_payment_events_payment_id ON payment_events(payment_id);
COMMIT;