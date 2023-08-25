-- Создание таблицы segments
CREATE TABLE segments (
                          id SERIAL PRIMARY KEY,
                          slug VARCHAR(255) NOT NULL
);

-- Создание таблицы users
CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       username VARCHAR(255) NOT NULL
);

-- Создание таблицы user_segments
CREATE TABLE user_segments (
                               user_id INT REFERENCES users(id),
                               segment_id INT REFERENCES segments(id),
                               PRIMARY KEY (user_id, segment_id)
);