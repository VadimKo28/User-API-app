package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *User) GetAll(c *gin.Context) {
  c.String(http.StatusOK, "За встречу\n")
}