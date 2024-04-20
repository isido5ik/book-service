CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    page_number INTEGER,
    date INTEGER,
    rating INTEGER
);
