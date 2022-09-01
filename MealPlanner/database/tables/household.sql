CREATE TABLE IF NOT EXISTS household (
    id SERIAL PRIMARY KEY,
    householdid VARCHAR(100) UNIQUE,
    grocerylist VARCHAR(100),
    plans int [],
    meals int []
);