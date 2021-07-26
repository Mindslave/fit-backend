-- name: CreateUser :one
INSERT INTO users (
  first_name, 
  last_name, 
  coach, 
  goal
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY name
LIMIT $1
OFFSET $2;