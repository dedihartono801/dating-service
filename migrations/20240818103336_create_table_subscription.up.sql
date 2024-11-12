CREATE TABLE "subscription" (
    id BIGSERIAL NOT NULL,
    user_id BIGINT NOT NULL,
    transaction_id BIGINT NOT NULL,
    package_type_id INT NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT FALSE,
    expires_at DATE NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY(id)
);