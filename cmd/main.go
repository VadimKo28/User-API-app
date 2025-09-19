package main

import (
	"api_app/internal/handlers/user"
	"github.com/gin-gonic/gin"
	
)



func main() {
	u := &user.User{}

	router := gin.New()

	user.RegisterEndpoints(router, u)

	router.Run(":3000")
}