package user

import (
	dto "api_app/internal/dto/user"
	"context"
	"database/sql"

	"github.com/lib/pq"
)

type Repository struct {
	db *sql.DB
}

func(r Repository) getAll(ctx context.Context) ([]dto.User, error){
  result , err := r.db.QueryContext(ctx, `select * from users;`)

	if err != nil {
		return nil, err
	}


	return result ,nil
}