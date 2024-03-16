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
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS files;
-- +goose StatementEnd
