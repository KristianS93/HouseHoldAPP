CREATE TABLE IF NOT EXISTS household (
    id SERIAL PRIMARY KEY,
    householdid VARCHAR(100) UNIQUE,
    meals int[],
    grocerylist VARCHAR(100)
);