SELECT e.id, name
FROM exercises_workouts ew
INNER JOIN exercise e ON e.id = ew.exercise_id
WHERE sc.workout_id = Y;