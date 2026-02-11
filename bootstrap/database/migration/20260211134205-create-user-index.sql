
-- +migrate Up
CREATE UNIQUE INDEX idx_users_email_unique ON "user"(email);

-- +migrate Down
DROP INDEX IF EXISTS idx_users_email_unique;
