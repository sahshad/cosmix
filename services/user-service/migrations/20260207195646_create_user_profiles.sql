-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_profiles (
    id SERIAL PRIMARY KEY,
    user_id UUID NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    username VARCHAR(255) UNIQUE,
    date_of_birth DATE,
    avatar_url TEXT,
    bio TEXT,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_profiles;
-- +goose StatementEnd
