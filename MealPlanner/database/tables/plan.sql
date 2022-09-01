CREATE TABLE IF NOT EXISTS plan(
    id SERIAL PRIMARY KEY,
    weekno int,
    householdid VARCHAR(100),
    meals int []
);