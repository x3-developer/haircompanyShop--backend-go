CREATE TABLE dashboard_users
(
    id          SERIAL PRIMARY KEY,
    email       VARCHAR(255) NOT NULL UNIQUE,
    password    VARCHAR(255) NOT NULL,
    role        VARCHAR(32)  NOT NULL,
    created_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Триггер для автоматического обновления updated_at
CREATE OR REPLACE FUNCTION set_updated_at_timestamp()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_dashboard_users_updated_at
    BEFORE UPDATE ON dashboard_users
    FOR EACH ROW
EXECUTE FUNCTION set_updated_at_timestamp();
