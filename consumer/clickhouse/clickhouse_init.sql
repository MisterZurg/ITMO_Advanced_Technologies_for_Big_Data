-- TODO: CHECK Fields and STuff
CREATE TABLE IF NOT EXISTS logs
(
    log_id              UInt32,
    timestamp           DateTime('Europe/Moscow'),
    level_of_msg        UInt8,
    application         String,
    hostname            String,
    logger_name         String,
    source              String,
    pid                 UInt32,
    line                String,
    error_status        UInt32,
    message             String
) Engine = Memory