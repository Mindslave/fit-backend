-- name: GetAllExercisesInWorkout :many
SELECT e.id, name
FROM exercises_workouts ew
INNER JOIN exercise e ON e.id = ew.exercise_id
WHERE ew.workout_id = $1;

-- name: GetAllWorkoutsWithExercise :many
select w.id
FROM exercises_workouts ew
INNER JOIN workouts w ON w.id = ew.workout_id
WHERE ew.workout_id = $1;