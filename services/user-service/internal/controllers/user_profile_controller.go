package controllers

import (
	"net/http"
	"strconv"

	"user-service/internal/dto"
	"user-service/internal/services"

	"github.com/gin-gonic/gin"
)

type UserProfileController struct {
	service services.UserProfileService
}

func NewUserProfileController(service services.UserProfileService) *UserProfileController {
	return &UserProfileController{service: service}
}

func (ctrl *UserProfileController) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "user service is ok"})
}

func (ctrl *UserProfileController) GetMe(c *gin.Context) {
	ctrl.GetMyProfile(c)
}

func (ctrl *UserProfileController) GetMyProfile(c *gin.Context) {
	userIDStr := c.GetHeader("X-User-Id")
	if userIDStr == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	profile, err := ctrl.service.GetProfile(uint(userID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, profile)
}

func (ctrl *UserProfileController) GetProfileByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	profile, err := ctrl.service.GetProfileByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, profile)
}

func (ctrl *UserProfileController) UpdateMe(c *gin.Context) {
	ctrl.UpdateMyProfile(c)
}

func (ctrl *UserProfileController) UpdateMyProfile(c *gin.Context) {
	userIDStr := c.GetHeader("X-User-Id")
	if userIDStr == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	var input dto.UpdateProfileDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	profile, err := ctrl.service.UpdateProfile(uint(userID), input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, profile)
}

func (ctrl *UserProfileController) GetByUsername(c *gin.Context) {
	username := c.Param("username")

	profile, err := ctrl.service.GetProfileByUsername(username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, profile)
}
