package user

import (
	dto "api_app/internal/dto/user"
	"context"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r Repository) GetAll(ctx context.Context) ([]dto.User, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT id, name, age FROM users`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []dto.User

	for rows.Next() {
		var user dto.User
		err := rows.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			log.Printf("Error scanning user: %v", err)
			return nil, err
		}

		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error during rows iteration: %v", err)
		return nil, err
	}

	return users, nil
}
