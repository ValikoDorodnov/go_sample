package repository

import (
	"context"

	"github.com/ValikoDorodnov/go_sample/internal/entity"
	"github.com/jmoiron/sqlx"
)

type GreetingRepository struct {
	db *sqlx.DB
}

func NewGreetingRepository(db *sqlx.DB) *GreetingRepository {
	return &GreetingRepository{
		db: db,
	}
}

func (r *GreetingRepository) All(ctx context.Context) ([]*entity.GreetingEntity, error) {
	var data []*entity.GreetingEntity

	query := `SELECT * FROM greeting`

	err := r.db.SelectContext(ctx, &data, query)
	if err != nil {
		return nil, err
	}

	return data, nil
}
