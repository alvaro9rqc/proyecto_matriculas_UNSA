-- +goose Up
-- +goose StatementBegin
/* 
  1. Add new fields for oauth and mistake with attendee_group_id
  2. Add foreing key to a provider id
  Add the attendee_group_id to correct table
*/
ALTER TABLE account_user 
DROP CONSTRAINT IF EXISTS account_user_attendee_group_id_fkey,
DROP COLUMN IF EXISTS attendee_group_id,
ADD COLUMN provider_user_id VARCHAR(255),
ADD COLUMN profile_picture TEXT,
ADD COLUMN oauth_provider_id SMALLINT;

ALTER TABLE account_user
ADD CONSTRAINT account_user_oauth_provider_id_fkey
FOREIGN KEY (oauth_provider_id) 
REFERENCES oauth_provider(id)
ON DELETE SET NULL;

ALTER TABLE attendee 
ADD COLUMN attendee_group_id SMALLINT NOT NULL;

ALTER TABLE attendee 
ADD CONSTRAINT attendee_attendee_group_id_fkey
FOREIGN KEY(attendee_group_id) 
REFERENCES attendee_group(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE attendee 
DROP CONSTRAINT IF EXISTS attendee_attendee_group_id_fkey;

ALTER TABLE attendee 
DROP COLUMN IF EXISTS attendee_group_id;

ALTER TABLE account_user
DROP CONSTRAINT IF EXISTS account_user_oauth_provider_id_fkey;

ALTER TABLE account_user 
DROP COLUMN IF EXISTS provider_user_id,
DROP COLUMN IF EXISTS profile_picture,
DROP COLUMN IF EXISTS oauth_provider_id,
ADD COLUMN attendee_group_id SMALLINT;

ALTER TABLE account_user
ADD CONSTRAINT account_user_attendee_group_id_fkey
FOREIGN KEY (attendee_group_id)
REFERENCES attendee_group(id)
ON DELETE RESTRICT;
-- +goose StatementEnd
