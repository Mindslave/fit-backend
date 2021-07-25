CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "coach" boolean NOT NULL,
  "goal" varchar NOT NULL
);

CREATE TABLE "workouts_users" (
  "workout_id" bigserial,
  "user_id" bigserial
);

CREATE TABLE "workouts" (
  "id" bigserial PRIMARY KEY,
  "time_estimate" integer
);

CREATE TABLE "exercise" (
  "id" bigserial PRIMARY KEY,
  "weight" integer,
  "reps" integer,
  "time_between_reps" integer,
  "explanation" varchar,
  "required_equipment" varchar,
  "muscle_area" varchar
);

CREATE TABLE "workouts_exercises" (
  "exercise_id" bigserial,
  "workout_id" bigserial
);

ALTER TABLE "workouts_users" ADD FOREIGN KEY ("workout_id") REFERENCES "workouts" ("id");

ALTER TABLE "workouts_users" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "workouts_exercises" ADD FOREIGN KEY ("exercise_id") REFERENCES "exercise" ("id");

ALTER TABLE "workouts_exercises" ADD FOREIGN KEY ("workout_id") REFERENCES "workouts" ("id");