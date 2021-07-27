ALTER TABLE exercises_workouts
RENAME TO workouts_exercises; 

ALTER TABLE exercise
DROP COLUMN name;