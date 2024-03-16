-- name: InsertFile :one
INSERT INTO files (
    id, 
    file_extension, 
    file_type, 
    file_name, 
    original_file_name, 
    file_size, 
    deletion_key
) VALUES (?, ?, ?, ?, ?, ?, ?) RETURNING *;