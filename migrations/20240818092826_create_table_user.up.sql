CREATE TYPE gender_enum AS ENUM ('male', 'female');
CREATE TABLE "user" (
    id BIGSERIAL NOT NULL,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(255) NOT NULL,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    gender gender_enum NOT NULL DEFAULT 'male',
    age INT NOT NULL,
    date_of_birth DATE NOT NULL,
    profile_picture VARCHAR(255),
    bio TEXT,
    location VARCHAR(255), 
    is_verified BOOLEAN NOT NULL DEFAULT FALSE,
    is_premium BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NULL,
    deleted_at TIMESTAMPTZ DEFAULT NULL,
    PRIMARY KEY(id)
);