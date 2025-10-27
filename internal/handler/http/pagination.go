package http

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/srgjo27/agora/internal/usecase"
)

const (
	defaultPage  = 1
	defaultLimit = 10
	maxLimit     = 100
)

func getPaginationParams(c *gin.Context) (usecase.PaginationParams, error) {
	page, err := strconv.Atoi(c.DefaultQuery("page", strconv.Itoa(defaultPage)))
	if err != nil || page < 1 {
		page = defaultPage
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", strconv.Itoa(defaultLimit)))
	if err != nil || limit <= 0 {
		limit = defaultLimit
	}

	if limit > maxLimit {
		limit = maxLimit
	}

	offset := (page - 1) * limit

	return usecase.PaginationParams{
		Limit:  limit,
		Offset: offset,
	}, nil
}
