CREATE TABLE IF NOT EXISTS users
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

INSERT INTO users 
    (name)  
VALUES 
    ('Kirill'),
    ('Marina');

CREATE TABLE IF NOT EXISTS segments
(
    id SERIAL PRIMARY KEY,
    slug VARCHAR(255) UNIQUE NOT NULL
);


CREATE TABLE IF NOT EXISTS users_segments 
(
    user_id INTEGER,
    segment_id INTEGER,

    PRIMARY KEY (user_id, segment_id)
);

ALTER TABLE users_segments ADD FOREIGN KEY (user_id) REFERENCES users (id);

ALTER TABLE users_segments ADD FOREIGN KEY (segment_id) REFERENCES segments (id);