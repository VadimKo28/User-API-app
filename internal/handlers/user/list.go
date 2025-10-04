package user

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *User) GetAll(c *gin.Context) {
	users, err := u.repo.GetAll(c.Request.Context())
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}
