DO $$
BEGIN
    CREATE TYPE LANGUAGE_STATUS AS ENUM ('in-progress', 'available', 'unavailable');
EXCEPTION
    WHEN duplicate_object THEN null;
END$$;

DO $$
BEGIN
    CREATE TYPE OFFER_TYPE AS ENUM ('%', 'x');
EXCEPTION
    WHEN duplicate_object THEN null;
END$$;

CREATE TABLE IF NOT EXISTS files (
    id SERIAL NOT NULL UNIQUE,

    path VARCHAR NOT NULL UNIQUE,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP,

    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS languages (
    id SERIAL NOT NULL UNIQUE,
    code CHAR(5) NOT NULL UNIQUE,
    status LANGUAGE_STATUS DEFAULT 'unavailable',

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP,

    PRIMARY KEY (id,code)
);

CREATE TABLE IF NOT EXISTS users (
    id SERIAL NOT NULL UNIQUE,
    language CHAR(5) DEFAULT 'en-EN',
    avatar_id INTEGER,

    username VARCHAR(25) NOT NULL UNIQUE,
    password VARCHAR(75) NOT NULL,
    email VARCHAR(60) NOT NULL UNIQUE,
    name VARCHAR(25),
    last_name VARCHAR(25),
    birthday TIMESTAMP,
    theme VARCHAR NOT NULL DEFAULT 'light',
    currency VARCHAR NOT NULL DEFAULT 'USD',

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP,

    PRIMARY KEY (id),
    FOREIGN KEY (language) REFERENCES languages(code) ON DELETE SET NULL,
    FOREIGN KEY (avatar_id) REFERENCES files(id) ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS products (
    id SERIAL NOT NULL UNIQUE,
    user_id INTEGER,

    name VARCHAR(80) NOT NULL,
    description VARCHAR(3000),
    price NUMERIC(20, 2) NOT NULL, 

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP,

    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS offers (
    id SERIAL NOT NULL UNIQUE,
    product_id INTEGER UNIQUE,

    type OFFER_TYPE NOT NULL,
    value VARCHAR(8) NOT NULL,
    description VARCHAR(2000),

    expirated_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP,

    PRIMARY KEY (id),
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS categories (
    id SERIAL NOT NULL UNIQUE,

    name VARCHAR(25) NOT NULL UNIQUE,
    description VARCHAR,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP,

    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS product_categories (
    id SERIAL NOT NULL UNIQUE,
    category_id INTEGER,
    product_id INTEGER,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS files_products (
    id SERIAL NOT NULL UNIQUE,
    file_id INTEGER,
    product_id INTEGER,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    PRIMARY KEY (id),
    FOREIGN KEY (file_id) REFERENCES files(id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);
