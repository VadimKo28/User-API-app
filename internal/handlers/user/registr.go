package user

import "github.com/gin-gonic/gin"

func RegisterEndpoints(r gin.IRouter, h *User) {
  r.GET("/users", h.GetAll) 
	r.POST("/users", h.CreateUser)
}

