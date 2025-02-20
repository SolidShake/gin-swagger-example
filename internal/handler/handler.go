package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

// @Summary      Get all users
// @Description  Returns list of all users
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {array}   User
// @Router       /users [get]
func (h *Handler) GetUsers(c *gin.Context) {
	users := []User{
		{ID: "1", Username: "user1", Email: "user1@example.com"},
		{ID: "2", Username: "user2", Email: "user2@example.com"},
	}
	c.JSON(http.StatusOK, users)
}
