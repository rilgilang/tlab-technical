-- +migrate Up
CREATE TABLE trx_log (
                        id UUID PRIMARY KEY,
                        sender UUID NOT NULL,
                        receiver UUID NOT NULL,
                        reason varchar NOT NULL,
                        status varchar NOT NULL,
                        amount bigint NOT NULL,
                        created_at TIMESTAMP,
                        updated_at TIMESTAMP


);

CREATE INDEX idx_trx_sender ON trx_log(sender);
CREATE INDEX idx_trx_receiver ON trx_log(receiver);
CREATE INDEX idx_wallet_id ON wallet(id);
CREATE INDEX idx_wallet_user_id ON wallet(user_id);

-- +migrate Down
DROP TABLE IF EXISTS trx_log;
DROP INDEX IF EXISTS idx_trx_sender;
DROP INDEX IF EXISTS idx_trx_receiver;
DROP INDEX IF EXISTS idx_wallet_id;
DROP INDEX IF EXISTS idx_wallet_user_id;
