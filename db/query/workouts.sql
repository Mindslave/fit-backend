-- name: CreateWorkout :one
INSERT INTO workouts (
  time_estimate 
) VALUES (
  $1
)
RETURNING *;

-- name: GetWorkout :one
SELECT * FROM workouts
WHERE id = $1 LIMIT 1;

-- name: ListWorkouts :many
SELECT * FROM workouts
ORDER BY name
LIMIT $1
OFFSET $2;