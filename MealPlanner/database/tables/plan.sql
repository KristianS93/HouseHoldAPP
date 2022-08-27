CREATE TABLE IF NOT EXISTS plan(
    id SERIAL PRIMARY KEY,
    weekno VARCHAR(100),
    householdid VARCHAR(100),
    meals int[]
);