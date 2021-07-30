package exercises

import "context"

type Exercise struct {
	id 		int
	name 	string
}

type ExerciseRepository interface {
	GetAll(ctx context.Context) (exercises []Exercise, err error)
	GetByID(ctx context.Context) (Exercise, error)
	GetByName(ctx context.Context) (Exercise, error)
	Update(ctx context.Context) error
	Store(ctx context.Context) error
	Delete(ctx context.Context) error
}