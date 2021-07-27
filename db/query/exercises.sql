-- name: CreateExercise :one
INSERT INTO exercise (
  name,
  weight, 
  reps, 
  time_between_reps, 
  explanation,
  required_equipment,
  muscle_area
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;

-- name: GetExercise :one
SELECT * FROM exercise
WHERE id = $1 LIMIT 1;

-- name: ListExercises :many
SELECT * FROM exercise
ORDER BY name
LIMIT $1
OFFSET $2;