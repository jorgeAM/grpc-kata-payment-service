BEGIN;

CREATE TABLE IF NOT EXISTS payment_schema.payments
(
    id uuid PRIMARY KEY,
    customer_id uuid NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'PENDING',
    order_id uuid NOT NULL,
    total_price DECIMAL(10,2) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ
);

COMMIT;
