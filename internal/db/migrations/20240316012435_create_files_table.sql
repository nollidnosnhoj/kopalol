-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS files (
  id TEXT NOT NULL,
  file_extension TEXT NOT NULL,
  file_type TEXT NOT NULL,
  file_name TEXT NOT NULL,
  original_file_name TEXT NOT NULL,
  file_size INTEGER NOT NULL,
  md5_hash TEXT NOT NULL,
  deletion_key TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);

CREATE INDEX idx_deletion_key ON files (deletion_key);
CREATE INDEX idx_created_at ON files (created_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_deletion_key;
DROP INDEX IF EXISTS idx_created_at;
DROP TABLE IF EXISTS files;
-- +goose StatementEnd
