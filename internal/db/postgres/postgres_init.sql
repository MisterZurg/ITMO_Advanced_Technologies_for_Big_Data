CREATE TABLE IF NOT EXISTS logs
(
    log_id       SERIAL PRIMARY KEY,
    timestamp    TIMESTAMP,
    level_of_msg INT,
    pid          INT,
    source       TEXT,
    error_status INT
);