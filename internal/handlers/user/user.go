package user

import (
	"api_app/internal/repository/user"

	_ "github.com/gin-gonic/gin"
)

type User struct {
	repo *user.Repository
}

func NewUser(repo *user.Repository) *User {
	return &User{repo: repo}
}

