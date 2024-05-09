-- +goose Up
-- +goose StatementBegin

CREATE TYPE oauth2_account_provider_enum AS ENUM
    ('facebook', 'instagram', 'twitter', 'google', 'linkedin');

CREATE TABLE IF NOT EXISTS users (
     id SERIAL PRIMARY KEY,
     name VARCHAR(255) NOT NULL,
     email VARCHAR(255) NOT NULL,
     password VARCHAR(255) NOT NULL,
     phone VARCHAR(255) DEFAULT NULL,
     photo_url VARCHAR(255) DEFAULT NULL,
     email_verified BOOLEAN DEFAULT FALSE,
     oauth2_account_provider oauth2_account_provider_enum DEFAULT NULL,
     created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS oauth2_states(
    id SERIAL PRIMARY KEY,
    state VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    valid_till TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS companies (
     id SERIAL PRIMARY KEY,
     name VARCHAR(255) NOT NULL,
     logo_url VARCHAR(255) DEFAULT NULL,
         description VARCHAR(255) DEFAULT NULL,
     url VARCHAR(255) DEFAULT NULL,
     user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
     created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS offers (
     id SERIAL PRIMARY KEY,
     company_id INTEGER REFERENCES companies(id) ON DELETE CASCADE,
     sale INTEGER NOT NULL,
     is_personal BOOLEAN DEFAULT true,
     created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
     expired_at TIMESTAMP DEFAULT NULL
);
CREATE TABLE IF NOT EXISTS users_offers (
      id SERIAL PRIMARY KEY,
      user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
      offer_id INTEGER REFERENCES offers(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS companies_administrators (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    company_id INTEGER REFERENCES companies(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS nonces (
      id SERIAL PRIMARY KEY,
      identifier VARCHAR(255) NOT NULL,
      nonce VARCHAR(255) NOT NULL,
      expired_at TIMESTAMP NOT NULL
);
CREATE TABLE IF NOT EXISTS offers_usages (
      id SERIAL PRIMARY KEY,
      user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
      offer_id INTEGER REFERENCES offers(id) ON DELETE CASCADE,
      created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
      expired_at TIMESTAMP DEFAULT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS users_offers;
DROP TABLE IF EXISTS offers_usages;
DROP TABLE IF EXISTS companies_administrators;
DROP TABLE IF EXISTS offers;
DROP TABLE IF EXISTS companies;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS nonces;

DROP TYPE IF EXISTS oauth2_account_provider_enum;
-- +goose StatementEnd
