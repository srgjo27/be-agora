package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/srgjo27/agora/internal/config"
	"github.com/srgjo27/agora/internal/domain"
	"github.com/srgjo27/agora/internal/usecase"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
	cfg         *config.Config
}

func NewUserHandler(uu usecase.UserUsecase, cfg *config.Config) *UserHandler {
	return &UserHandler{userUsecase: uu, cfg: cfg}
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
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}

		return
	}

	c.JSON(http.StatusCreated, NewUserResponse(user))
}

func (h *UserHandler) Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, refreshToken, err := h.userUsecase.Login(c.Request.Context(), req.Email, req.Password)

	if err != nil {
		switch err {
		case domain.ErrUnauthorized:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
		case domain.ErrInvalid:
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}

		return
	}

	c.SetCookie("refresh_token", refreshToken, int(h.cfg.RefreshTokenDurationHours*3600), "/api/v1/auth", h.cfg.CookieDomain, h.cfg.CookieSecure, true)

	c.JSON(http.StatusOK, LoginResponse{
		AccessToken: accessToken,
	})
}

func (h *UserHandler) Refresh(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "refresh token not found"})
		return
	}

	newAccessToken, err := h.userUsecase.Refresh(c.Request.Context(), refreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid refresh token"})
		return
	}

	c.JSON(http.StatusOK, LoginResponse{
		AccessToken: newAccessToken,
	})
}

func (h *UserHandler) Logout(c *gin.Context) {
	c.SetCookie(
		"refresh_token",
		"",
		-1,
		"/api/v1/auth",
		h.cfg.CookieDomain,
		h.cfg.CookieSecure,
		true,
	)

	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}
