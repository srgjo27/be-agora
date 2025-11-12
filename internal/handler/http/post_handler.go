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

type PostHandler struct {
	postUsecase usecase.PostUsecase
}

func NewPostHandler(pu usecase.PostUsecase) *PostHandler {
	return &PostHandler{postUsecase: pu}
}

func getThreadIDFromParam(c *gin.Context) (uuid.UUID, error) {
	idParam := c.Param("thread_id")
	threadID, err := uuid.Parse(idParam)
	return threadID, err
}

func (h *PostHandler) Create(c *gin.Context) {
	threadID, err := getThreadIDFromParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid thread ID"})

		return
	}

	var req CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	userID, exists := getUserIDFromCtx(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})

		return
	}

	post, err := h.postUsecase.Create(c.Request.Context(), req.Content, userID, threadID, req.ParentPostID)
	if err != nil {
		if err == domain.ErrInvalid {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input or thread ID"})

			return
		}

		if err == domain.ErrThreadLocked {
			c.JSON(http.StatusForbidden, gin.H{"error": "thread is locked"})

			return
		}

		log.Fatalf("[ERROR]: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})

		return
	}

	c.JSON(http.StatusOK, NewPostResponse(post, nil))
}

func (h *PostHandler) GetByThreadID(c *gin.Context) {
	threadID, err := getThreadIDFromParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid thread ID"})

		return
	}

	params, err := getPaginationParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid pagination parameters"})

		return
	}

	posts, userMap, totalItems, err := h.postUsecase.GetByThreadID(c.Request.Context(), threadID, params)
	if err != nil {
		if err == domain.ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "thread not found"})

			return
		}

		log.Fatalf("[ERROR]: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})

		return
	}

	dtos := make([]*PostResponse, len(posts))
	for i, p := range posts {
		dtos[i] = NewPostResponse(p, userMap[p.UserID])
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

	response := gin.H{
		"data": dtos,
		"meta": meta,
	}

	c.JSON(http.StatusOK, response)
}
