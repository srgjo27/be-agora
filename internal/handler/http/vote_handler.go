package http

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/srgjo27/agora/internal/domain"
	"github.com/srgjo27/agora/internal/usecase"
)

type VoteHandler struct {
	voteUsecase usecase.VoteUsecase
}

func NewVoteHandler(vu usecase.VoteUsecase) *VoteHandler {
	return &VoteHandler{voteUsecase: vu}
}

func (h *VoteHandler) VoteOnThread(c *gin.Context) {
	idParam := c.Param("thread_id")
	threadID, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid thread ID"})

		return
	}

	var req VoteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid vote_type, must be 1, 0, or -1"})

		return
	}

	userID, exists := getUserIDFromCtx(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})

		return
	}

	err = h.voteUsecase.VoteOnThread(c.Request.Context(), userID, threadID, req.VoteType)
	if err != nil {
		if err == domain.ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "thread not found"})

			return
		}

		log.Fatalf("[ERROR]: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})

		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "vote recorded"})
}
