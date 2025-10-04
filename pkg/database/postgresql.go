package database

import (
	_"github.com/lib/pq"
	"database/sql"
	"fmt"
)

type ConnectionInfo struct {
	Host string 
	Port string 
	Username string 
	DBName string 
	SSLMode string 
	Password string
}

func NewPostgresConnection(info ConnectionInfo) (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
	info.Username, info.Password, info.Host, info.Port, info.DBName))

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Connection")

	return db, nil
}
