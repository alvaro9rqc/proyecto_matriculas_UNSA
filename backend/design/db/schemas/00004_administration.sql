-- +goose Up
-- +goose StatementBegin

-- Institution Table: Represents a university or educational institution
CREATE TABLE institution (
    id SERIAL PRIMARY KEY,
    name VARCHAR(128) UNIQUE NOT NULL,
    logo_url VARCHAR(255)  -- Optional logo/avatar URL
);

CREATE TABLE administrative (
    id SERIAL PRIMARY KEY,
    account_id INTEGER UNIQUE NOT NULL,
    process_id INTEGER NOT NULL,
    FOREIGN KEY (account_id) REFERENCES account (id) ON DELETE CASCADE,
    FOREIGN KEY (process_id) REFERENCES process (id) ON DELETE RESTRICT
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS admin_user;

DROP TABLE IF EXISTS institution;

-- +goose StatementEnd
