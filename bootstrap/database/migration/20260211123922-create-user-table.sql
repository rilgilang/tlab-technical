-- +migrate Up
CREATE TABLE "user" (
                             id UUID PRIMARY KEY,
                             name VARCHAR NOT NULL,
                             email VARCHAR NOT NULL,
                             password VARCHAR NOT NULL,
                             created_at TIMESTAMP,
                             updated_at TIMESTAMP
);

-- +migrate Down
DROP TABLE IF EXISTS "user";