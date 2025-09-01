package handlers

import (
	"fmt"
	"tasks-backend/internal/repositories"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserRepo repositories.IUserRepository
}

func NewUserHandler(userRepo repositories.IUserRepository) *UserHandler {
	return &UserHandler{
		UserRepo: userRepo,
	}
}

func (h *UserHandler) RegisterRoutes(c *gin.RouterGroup) {
	c.POST("/", h.CreateUser)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	body := c.Request.Body

	fmt.Println("Printing body", body)
}
