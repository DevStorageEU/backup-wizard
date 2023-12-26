-- +goose Up
-- +goose StatementBegin
CREATE TABLE devices (
    id VARCHAR(36) PRIMARY KEY NOT NULL DEFAULT UUID(),
    name VARCHAR(128) NOT NULL,
    kind VARCHAR(16) NOT NULL,
    protection VARCHAR(16) NOT NULL,
    last_backup DATETIME,
    ips VARCHAR(1024) NOT NULL,
    agent VARCHAR(128) NOT NULL,
    created_at DATETIME NOT NULL DEFAULT NOW(),
    updated_at DATETIME NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE devices;
-- +goose StatementEnd
