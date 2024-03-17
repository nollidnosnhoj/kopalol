-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS files (
  id varchar NOT NULL,
  file_extension varchar NOT NULL,
  file_type varchar NOT NULL,
  file_name varchar NOT NULL,
  original_file_name varchar NOT NULL,
  file_size integer NOT NULL,
  deletion_key varchar NOT NULL,
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
