CREATE TABLE "package_type" (
    id BIGSERIAL NOT NULL,
    name VARCHAR(100) NOT NULL,
    price INT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NULL,
    PRIMARY KEY(id)
);