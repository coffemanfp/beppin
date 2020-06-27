DO $$
BEGIN
    CREATE TYPE LANGUAGE_STATUS AS ENUM ('in-progress', 'available', 'unavailable');
EXCEPTION
    WHEN duplicate_object THEN null;
END$$;

CREATE TABLE IF NOT EXISTS offers (
    id SERIAL NOT NULL UNIQUE,
    product_id INTEGER UNIQUE,

    type OFFER_TYPE NOT NULL,
    value VARCHAR,
    expirated_at TIMESTAMP,

    PRIMARY KEY (id),
    FOREIGN KEY (id) REFERENCES products(id)
);

CREATE TABLE IF NOT EXISTS languages (
    id SERIAL,

    code VARCHAR(4),
    status LANGUAGES_STATUS DEFAULT 'unavailable',

    created_at TIMESTAMP NOT NULl DEFAULT NOW(),
    updated_at TIMESTAMP,

    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS users (
    id SERIAL NOT NULL UNIQUE,
    language_id INTEGER,

    username VARCHAR UNIQUE,
    password VARCHAR,
    name VARCHAR NOT NULL,
    last_name VARCHAR NOT NULL,
    birthday TIMESTAMP NOT NULL,
    theme VARCHAR DEFAULT 'light',

    PRIMARY KEY (id),
    FOREIGN KEY (language_id) REFERENCES languages(id)
);

CREATE TABLE IF NOT EXISTS products (
    id SERIAL,
    user_id INTEGER,

    name VARCHAR NOT NULL,
    description VARCHAR,

    created_at TIMESTAMP NOT NULl DEFAULT NOW(),
    updated_at TIMESTAMP,

    PRIMARY KEY (id),
    FOREIGN KEY (id) REFERENCES offers(id),
    FOREIGN KEY (id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS categories (
    id SERIAL NOT NULL UNIQUE,

    name VARCHAR UNIQUE,
    related_categories VARCHAR[],

    PRIMARY KEY (id)
);