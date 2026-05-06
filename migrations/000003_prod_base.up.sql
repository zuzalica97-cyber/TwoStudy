CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    ProductName VARCHAR(255) NOT NULL,
    ProductDescription VARCHAR(255) NOT NULL,
    ProductCost INTEGER NOT NULL,
    ProductAmount INTEGER NOT NULL,
    ProductIDP INTEGER NOT NULL
);