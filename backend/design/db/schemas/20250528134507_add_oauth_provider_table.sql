-- +goose Up
-- +goose StatementBegin
CREATE TABLE oauth_provider (
  id SMALLSERIAL PRIMARY KEY,
  name VARCHAR(50) NOT NULL UNIQUE
);
-- +goose StatementEnd
