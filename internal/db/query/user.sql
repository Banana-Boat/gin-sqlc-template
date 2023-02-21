-- name: CreateUser :execresult
INSERT INTO users (
  username, password, gender, age
) VALUES (
  ?, ?, ?, ?
);

-- name: GetUser :one
SELECT * FROM users
WHERE id = ? LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id;

-- name: UpdateUser :exec
UPDATE users SET password = ?, gender = ?, age = ?
WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?;