-- +goose Up
-- +goose StatementBegin
CREATE TABLE devices (
    id uuid PRIMARY KEY NOT NULL,
    name VARCHAR(128) NOT NULL,
    hostname VARCHAR(128) NOT NULL,
    kind VARCHAR(16) NOT NULL,
    protection VARCHAR(16) NOT NULL,
    last_backup TIMESTAMP,
    cpu VARCHAR(128) NOT NULL,
    ram VARCHAR(16) NOT NULL,
    disks VARCHAR(32)[] NOT NULL,
    ips VARCHAR(64)[] NOT NULL,
    agent VARCHAR(128) NOT NULL,
    operating_system VARCHAR(128) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE devices;
-- +goose StatementEnd
