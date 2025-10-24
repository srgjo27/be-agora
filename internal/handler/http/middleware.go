package http

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/srgjo27/agora/internal/usecase"
)

const (
	userCtxKey = "userID"
	roleCtxKey = "userRole"
)

type AuthMiddleware struct {
	tokenSvc usecase.TokenService
}

func NewAuthMiddleware(ts usecase.TokenService) *AuthMiddleware {
	return &AuthMiddleware{tokenSvc: ts}
}

func (m *AuthMiddleware) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization header is required"})

			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header format"})

			return
		}

		tokenString := parts[1]

		userID, role, err := m.tokenSvc.ValidateToken(c.Request.Context(), tokenString)
		if err != nil {
			log.Fatalf("[ERROR]: %v", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})

			return
		}

		c.Set(userCtxKey, userID)
		c.Set(roleCtxKey, role)
	}
}

func getUserIDFromCtx(ctx *gin.Context) (uuid.UUID, bool) {
	val, ok := ctx.Get(userCtxKey)
	if !ok {
		return uuid.Nil, false
	}

	id, ok := val.(uuid.UUID)

	return id, ok
}

func getUserRoleFromCtx(ctx *gin.Context) (string, bool) {
	val, ok := ctx.Get(roleCtxKey)
	if !ok {
		return "", false
	}

	role, ok := val.(string)

	return role, ok
}

func (m *AuthMiddleware) AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := getUserRoleFromCtx(c)
		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "role not found in context"})

			return
		}

		if role != "admin" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "admin access required"})

			return
		}

		c.Next()
	}
}
