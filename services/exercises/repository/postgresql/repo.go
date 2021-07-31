package postgresql

import (
	"context"
	"database/sql"

	"github.com/Mindslave/fit-backend/services/exercises"
)

type PostgresExerciseRepository struct {
	DB *sql.DB
}

func NewPostgresExerciseRepository(db *sql.DB) exercises.ExerciseRepository {
	return &PostgresExerciseRepository {
		DB: db,
	}
}

func (per *PostgresExerciseRepository) GetAll(ctx context.Context) (excs []exercises.Exercise, err error) {

}

func (per *PostgresExerciseRepository) GetByID(ctx context.Context) (exc exercises.Exercise, err error) {

}
func (per *PostgresExerciseRepository) GetByName(ctx context.Context) (exc exercises.Exercise, err error){

}
func (per *PostgresExerciseRepository) Update(ctx context.Context) error {

}
func (per *PostgresExerciseRepository) Store(ctx context.Context) error {

}
func (per *PostgresExerciseRepository) Delete(ctx context.Context) error {

}