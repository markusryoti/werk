CREATE TABLE IF NOT EXISTS workout (
    id SERIAL PRIMARY KEY,
    date timestamp DEFAULT now(),
    name VARCHAR(255),
    uid CHAR(36)
);

CREATE TABLE IF NOT EXISTS movement (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    workout_id INT REFERENCES workout(id) ON DELETE CASCADE,
    uid CHAR(36)
);

CREATE TABLE IF NOT EXISTS movement_set (
    id SERIAL PRIMARY KEY,
    reps SMALLINT,
    weight SMALLINT,
    movement_id INT REFERENCES movement(id) ON DELETE CASCADE,
    uid CHAR(36)
);
