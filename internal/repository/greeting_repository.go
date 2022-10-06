package repository

import (
	"database/sql"
	"github.com/ValikoDorodnov/go_sample/internal/entity"
)

type GreetingRepository struct {
	db *sql.DB
}

func NewGreetingRepository(db *sql.DB) *GreetingRepository {
	return &GreetingRepository{
		db: db,
	}
}

func (r *GreetingRepository) All() ([]entity.GreetingEntity, error) {
	rows, err := r.db.Query("SELECT * FROM greeting")
	if err != nil {
		return nil, err
	}
	var data []entity.GreetingEntity
	for rows.Next() {
		var r entity.GreetingEntity
		err = rows.Scan(&r.Id, &r.Name)
		if err != nil {
			return nil, err
		}
		data = append(data, r)
	}

	return data, nil
}
