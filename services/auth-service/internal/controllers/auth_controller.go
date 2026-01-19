package controllers

import (
	"net/http"
	"os"

	"auth-service/internal/dto"
	"auth-service/internal/services"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (ctrl *AuthController) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "auth service is ok"})
}

func (ctrl *AuthController) Register(c *gin.Context) {
	var registerDTO dto.RegisterDTO
	if err := c.ShouldBindJSON(&registerDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := ctrl.authService.Register(registerDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user": gin.H{
			"id":         user.ID,
			"first_name": user.FirstName,
			"last_name":  user.LastName,
			"email":      user.Email,
			"user_name":  user.Username,
			"role":       user.Role,
		},
	})
}

func (ctrl *AuthController) Login(c *gin.Context) {
	var input dto.LoginDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	access, refresh, user, err := ctrl.authService.Login(input)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	secure := false
	domain := ""
	if os.Getenv("ENV") == "production" {
		secure = true
		domain = os.Getenv("COOKIE_DOMAIN")
	}

	c.SetCookie("refresh_token", refresh, 60*60*24*30, "/", domain, secure, true)

	c.JSON(http.StatusOK, gin.H{
		"accessToken": access,
		"user": gin.H{
			"id":         user.ID,
			"first_name": user.FirstName,
			"last_name":  user.LastName,
			"email":      user.Email,
			"user_name":  user.Username,
			"role":       user.Role,
		},
	})
}

func (ctrl *AuthController) Refresh(c *gin.Context) {
	rt, err := c.Cookie("refresh_token")
	if err != nil || rt == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no refresh token"})
		return
	}

	newAccess, newRefresh, err := ctrl.authService.Refresh(rt)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	secure := false
	domain := ""
	if os.Getenv("ENV") == "production" {
		secure = true
		domain = os.Getenv("COOKIE_DOMAIN")
	}
	c.SetCookie("refresh_token", newRefresh, 60*60*24*30, "/", domain, secure, true)

	c.JSON(http.StatusOK, gin.H{"accessToken": newAccess})
}

func (ctrl *AuthController) Logout(c *gin.Context) {
	c.SetCookie("refresh_token", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "logged out"})
}