DO $$
BEGIN
    CREATE TYPE LANGUAGE_STATUS AS ENUM ('in-progress', 'available', 'unavailable');
EXCEPTION
    WHEN duplicate_object THEN null;
END$$;

DO $$
BEGIN
    CREATE TYPE OFFER_TYPE AS ENUM ('percentage', 'promotion');
EXCEPTION
    WHEN duplicate_object THEN null;
END$$;

CREATE TABLE IF NOT EXISTS languages (
    id SERIAL,

    code VARCHAR(5),
    status LANGUAGE_STATUS DEFAULT 'unavailable',

    created_at TIMESTAMP NOT NULl DEFAULT NOW(),
    updated_at TIMESTAMP,

    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS users (
    id SERIAL NOT NULL UNIQUE,
    language_id INTEGER,

    username VARCHAR(25) UNIQUE,
    password VARCHAR(75),
    name VARCHAR(25) NOT NULL,
    last_name VARCHAR(25) NOT NULL,
    birthday TIMESTAMP NOT NULL,
    theme VARCHAR DEFAULT 'light',

    PRIMARY KEY (id),
    FOREIGN KEY (language_id) REFERENCES languages(id)
);

CREATE TABLE IF NOT EXISTS products (
    id SERIAL NOT NULL UNIQUE,
    user_id INTEGER,

    name VARCHAR(80) NOT NULL,
    description VARCHAR(3000),
    categories VARCHAR[],

    created_at TIMESTAMP NOT NULl DEFAULT NOW(),
    updated_at TIMESTAMP,

    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS offers (
    id SERIAL NOT NULL UNIQUE,
    product_id INTEGER UNIQUE,

    type OFFER_TYPE NOT NULL,
    value VARCHAR(8),
    description VARCHAR(2000),

    expirated_at TIMESTAMP,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,

    PRIMARY KEY (id),
    FOREIGN KEY (id) REFERENCES products(id)
);

CREATE TABLE IF NOT EXISTS categories (
    id SERIAL NOT NULL UNIQUE,

    name VARCHAR(25) UNIQUE,
    related_categories VARCHAR[],

    PRIMARY KEY (id)
);