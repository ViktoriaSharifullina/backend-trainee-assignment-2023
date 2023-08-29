-- Создание таблицы segments
CREATE TABLE segments (
                          id SERIAL PRIMARY KEY,
                          slug VARCHAR(255) NOT NULL UNIQUE,
                          auto_assign_percent INT
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
                               PRIMARY KEY (user_id, segment_id),
                               expires_at TIMESTAMP
);

-- Создание таблицы user_segment_history
CREATE TABLE user_segment_histories (
                                      id SERIAL PRIMARY KEY,
                                      user_id INT NOT NULL,
                                      segment_id INT NOT NULL,
                                      operation VARCHAR(10) NOT NULL,
                                      date TIMESTAMP NOT NULL,
                                      FOREIGN KEY (user_id) REFERENCES users(id),
                                      FOREIGN KEY (segment_id) REFERENCES segments(id)
);