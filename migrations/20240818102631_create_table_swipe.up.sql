CREATE TYPE swipe_enum AS ENUM ('like', 'pass');
CREATE TABLE "swipe" (
    id BIGSERIAL NOT NULL,
    swipper_user_id BIGINT NOT NULL,
    target_user_id BIGINT NOT NULL,
    swipe_type swipe_enum NOT NULL DEFAULT 'like',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY(id)
);