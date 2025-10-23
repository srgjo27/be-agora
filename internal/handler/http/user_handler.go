package http

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/srgjo27/agora/internal/domain"
	"github.com/srgjo27/agora/internal/usecase"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(uu usecase.UserUsecase) *UserHandler {
	return &UserHandler{userUsecase: uu}
}

func (h *UserHandler) Register(c *gin.Context) {
	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	user, err := h.userUsecase.Register(c.Request.Context(), req.Username, req.Email, req.Password)

	if err != nil {
		switch err {
		case domain.ErrConflict:
			c.JSON(http.StatusConflict, gin.H{"error": "email already exists"})
		case domain.ErrInvalid:
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		default:
			log.Printf("ERROR: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}

		return
	}

	c.JSON(http.StatusCreated, NewUserResponse(user))
}
