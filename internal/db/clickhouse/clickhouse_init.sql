CREATE TABLE IF NOT EXISTS logs
(
    log_id       SERIAL PRIMARY KEY,
    timestamp    TIMESTAMP,
    level_of_msg INT,
    application  TEXT,
    hostname     TEXT,
    logger       name,
    source       TEXT,
    pid          INT,
    error_status INT,
    message      TEXT
);