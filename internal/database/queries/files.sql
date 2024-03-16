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

-- name: GetFileForDeletion :one
SELECT * FROM files WHERE id = ? AND deletion_key = ? LIMIT 1;

-- name: DeleteFile :exec
DELETE FROM files WHERE id = ?;