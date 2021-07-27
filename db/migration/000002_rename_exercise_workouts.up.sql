ALTER TABLE workouts_exercises 
RENAME TO exercises_workouts; 

ALTER TABLE exercise
ADD COLUMN name varchar;