-- Удаляем \connect family (не нужно, скрипт выполняется в новой БД)
CREATE TABLE users (
    id SERIAL PRIMARY KEY,  -- SERIAL вместо integer + PRIMARY KEY
    username TEXT NOT NULL,
    password TEXT NOT NULL
);

CREATE TABLE category (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    user_id INTEGER NOT NULL,
    created_at DATE DEFAULT CURRENT_DATE NOT NULL
);

CREATE TABLE expense (
    id SERIAL PRIMARY KEY,
    category_id INTEGER NOT NULL,
    amount NUMERIC(10,2) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    description TEXT,
    user_id INTEGER
);