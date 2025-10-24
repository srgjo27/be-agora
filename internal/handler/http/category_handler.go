package http

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/srgjo27/agora/internal/domain"
	"github.com/srgjo27/agora/internal/usecase"
)

type CategoryHandler struct {
	categoryUsecase usecase.CategoryUsecase
}

func NewCategoryHandler(cu usecase.CategoryUsecase) *CategoryHandler {
	return &CategoryHandler{categoryUsecase: cu}
}

func (h *CategoryHandler) Create(c *gin.Context) {
	var req CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	cat, err := h.categoryUsecase.Create(c.Request.Context(), req.Name, req.Description)
	if err != nil {
		if err == domain.ErrConflict {
			c.JSON(http.StatusConflict, gin.H{"error": "category already exists"})

			return
		}

		log.Fatalf("[ERROR]: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})

		return
	}

	c.JSON(http.StatusOK, NewCategoryResponse(cat))
}

func (h *CategoryHandler) GetAll(c *gin.Context) {
	cats, err := h.categoryUsecase.GetAll(c.Request.Context())
	if err != nil {
		log.Fatalf("[ERROR]: %v", err)

		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})

		return
	}

	c.JSON(http.StatusOK, NewCategoryListResponse(cats))
}
