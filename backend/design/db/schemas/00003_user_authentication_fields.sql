-- +goose Up
-- +goose StatementBegin
/*
1. Agrega nuevos campos para oauth y corrige el uso de student_group_id
2. Agrega la clave foránea a provider id
3. Añade student_group_id a la tabla correcta
*/
ALTER TABLE account_user
DROP CONSTRAINT IF EXISTS account_user_attendee_group_id_fkey,
DROP COLUMN IF EXISTS attendee_group_id,
ADD COLUMN provider_user_id VARCHAR(255),
ADD COLUMN profile_picture TEXT,
ADD COLUMN oauth_provider_id SMALLINT;

ALTER TABLE account_user ADD CONSTRAINT account_user_oauth_provider_id_fkey FOREIGN KEY (oauth_provider_id) REFERENCES oauth_provider (id) ON DELETE SET NULL;

ALTER TABLE student
ADD COLUMN student_group_id SMALLINT NOT NULL;

ALTER TABLE student ADD CONSTRAINT student_student_group_id_fkey FOREIGN KEY (student_group_id) REFERENCES student_group (id);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
ALTER TABLE student
DROP CONSTRAINT IF EXISTS student_student_group_id_fkey;

ALTER TABLE student
DROP COLUMN IF EXISTS student_group_id;

ALTER TABLE account_user
DROP CONSTRAINT IF EXISTS account_user_oauth_provider_id_fkey;

ALTER TABLE account_user
DROP COLUMN IF EXISTS provider_user_id,
DROP COLUMN IF EXISTS profile_picture,
DROP COLUMN IF EXISTS oauth_provider_id,
ADD COLUMN attendee_group_id SMALLINT;

ALTER TABLE account_user ADD CONSTRAINT account_user_attendee_group_id_fkey FOREIGN KEY (attendee_group_id) REFERENCES student_group (id) ON DELETE RESTRICT;

-- +goose StatementEnd
