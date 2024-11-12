CREATE TYPE status_enum AS ENUM ('unpaid', 'paid');
CREATE TABLE "transaction" (
    id BIGSERIAL NOT NULL,
    user_id BIGINT NOT NULL,
    payment_method_id INT NOT NULL,
    currency VARCHAR(3) NOT NULL,
    amount INT NOT NULL,
    status status_enum NOT NULL DEFAULT 'unpaid',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY(id)
);