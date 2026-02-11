-- +migrate Up
CREATE TABLE wallet (
                        id UUID PRIMARY KEY,
                        user_id UUID NOT NULL,
                        amount bigint NOT NULL,
                        created_at TIMESTAMP,
                        updated_at TIMESTAMP
);

CREATE UNIQUE INDEX idx_users_wallet_id_unique ON "wallet"(id);
CREATE UNIQUE INDEX idx_users_wallet_user_id_unique ON "wallet"(user_id);

-- +migrate Down
DROP TABLE IF EXISTS wallet;
DROP INDEX IF EXISTS idx_users_wallet_id_unique;
DROP INDEX IF EXISTS idx_users_wallet_user_id_unique;
