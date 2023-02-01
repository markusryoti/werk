CREATE TABLE IF NOT EXISTS workout (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    user_id VARCHAR(36),
    created_at timestamptz DEFAULT now(),
    updated_at timestamptz DEFAULT now()
);

CREATE TABLE IF NOT EXISTS movement (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    user_id VARCHAR(36),
    created_at timestamptz DEFAULT now(),
    updated_at timestamptz DEFAULT now()
);

CREATE TABLE IF NOT EXISTS workout_movement (
    id SERIAL PRIMARY KEY,
    workout_id INT REFERENCES workout(id) ON DELETE CASCADE,
    movement_id INT REFERENCES movement(id) ON DELETE CASCADE,
    user_id VARCHAR(36),
    created_at timestamptz DEFAULT now(),
    updated_at timestamptz DEFAULT now()
);

CREATE TABLE IF NOT EXISTS movement_set (
    id SERIAL PRIMARY KEY,
    reps SMALLINT,
    weight INT,
    workout_movement_id INT REFERENCES workout_movement(id) ON DELETE CASCADE,
    user_id VARCHAR(36),
    created_at timestamptz DEFAULT now(),
    updated_at timestamptz DEFAULT now()
);
