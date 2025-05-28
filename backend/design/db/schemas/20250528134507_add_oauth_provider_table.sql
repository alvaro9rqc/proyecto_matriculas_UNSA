-- +goose Up
-- +goose StatementBegin
CREATE TABLE auth_provider (
  id SMALLSERIAL PRIMARY KEY,
  name VARCHAR(50) NOT NULL UNIQUE
);
-- +goose StatementEnd
