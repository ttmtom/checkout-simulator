BEGIN;
    DROP TABLE IF EXISTS payment_events;
    DROP TABLE IF EXISTS payments;
    DROP TABLE IF EXISTS orders;
    DROP TYPE IF EXISTS payment_status_enum;
    DROP TYPE IF EXISTS order_status_enum;
COMMIT;