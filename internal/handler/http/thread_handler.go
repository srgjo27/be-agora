package http

import (
	"log"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/srgjo27/agora/internal/domain"
	"github.com/srgjo27/agora/internal/usecase"
)

type ThreadHandler struct {
	threadUsecase usecase.ThreadUsecase
}

func NewThreadHandler(tu usecase.ThreadUsecase) *ThreadHandler {
	return &ThreadHandler{threadUsecase: tu}
}

func (h *ThreadHandler) Create(c *gin.Context) {
	var req CreateThreadRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	userID, exists := getUserIDFromCtx(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})

		return
	}

	thread, err := h.threadUsecase.Create(c.Request.Context(), req.Title, req.Content, userID, req.CategoryID)
	if err != nil {
		if err == domain.ErrInvalid {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input or category ID"})

			return
		}
		log.Fatalf("[ERROR]: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})

		return
	}

	c.JSON(http.StatusCreated, NewThreadDetailResponse(thread))
}

func (h *ThreadHandler) GetAll(c *gin.Context) {
	params, err := getPaginationParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid pagination parameters"})

		return
	}

	threads, totalItems, err := h.threadUsecase.GetAll(c.Request.Context(), params)
	if err != nil {
		log.Fatalf("[ERROR]: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})

		return
	}

	totalPages := 0
	if params.Limit > 0 {
		totalPages = int(math.Ceil(float64(totalItems) / float64(params.Limit)))
	}

	meta := PaginationMeta{
		TotalItems:  totalItems,
		TotalPages:  totalPages,
		CurrentPage: (params.Offset / params.Limit) + 1,
		Limit:       params.Limit,
	}

	response := PaginatedThreadsResponse{
		Data: NewThreadListResponse(threads),
		Meta: meta,
	}

	c.JSON(http.StatusOK, response)
}

func (h *ThreadHandler) GetByID(c *gin.Context) {
	idParam := c.Param("thread_id")
	threadID, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid thread ID"})

		return
	}

	thread, err := h.threadUsecase.GetByID(c.Request.Context(), threadID)
	if err != nil {
		if err == domain.ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "thread not found"})

			return
		}

		log.Fatalf("[ERROR]: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})

		return
	}

	c.JSON(http.StatusOK, NewThreadDetailResponse(thread))
}
