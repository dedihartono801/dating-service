CREATE TABLE "match" (
    id BIGSERIAL NOT NULL,
    user_id BIGINT NOT NULL,
    matched_user_id BIGINT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY(id)
);