package postgresql

import (
	"context"
	"database/sql"
)

type store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *store {
	store := &store{
		db: db,
		Queries: New(db),
	}
	return store
}

//executes a function within a database transcation
func (s store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
}