package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *User) CreateUser(c *gin.Context) {
  c.String(http.StatusOK, "За встречу создано\n")
}