CREATE TABLE IF NOT EXISTS meal(
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    items int[]
);