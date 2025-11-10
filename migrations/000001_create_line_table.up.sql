CREATE TABLE lines
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(255) NOT NULL UNIQUE,
    color       VARCHAR(25)  NOT NULL DEFAULT '#000000',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Триггер для автоматического обновления updated_at
CREATE OR REPLACE FUNCTION set_updated_at_timestamp()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_lines_updated_at
    BEFORE UPDATE ON lines
    FOR EACH ROW
EXECUTE FUNCTION set_updated_at_timestamp();
