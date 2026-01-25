package repositories

import (
	"user-service/internal/models"

	"gorm.io/gorm"
)

type UserProfileRepository interface {
	Create(profile *models.UserProfile) error
	FindByUserID(userID uint) (*models.UserProfile, error)
	FindByID(id uint) (*models.UserProfile, error)
	FindByUsername(username string) (*models.UserProfile, error)
	Update(profile *models.UserProfile) error
	Delete(id uint) error
}

type userProfileRepo struct {
	db *gorm.DB
}

func NewUserProfileRepository(db *gorm.DB) UserProfileRepository {
	return &userProfileRepo{db: db}
}

func (repo *userProfileRepo) Create(profile *models.UserProfile) error {
	return repo.db.Create(profile).Error
}

func (repo *userProfileRepo) FindByUserID(userID uint) (*models.UserProfile, error) {
	var profile models.UserProfile
	if err := repo.db.Where("user_id = ?", userID).First(&profile).Error; err != nil {
		return nil, err
	}
	return &profile, nil
}

func (repo *userProfileRepo) FindByID(id uint) (*models.UserProfile, error) {
	var profile models.UserProfile
	if err := repo.db.First(&profile, id).Error; err != nil {
		return nil, err
	}
	return &profile, nil
}

func (repo *userProfileRepo) Update(profile *models.UserProfile) error {
	return repo.db.Save(profile).Error
}

func (repo *userProfileRepo) FindByUsername(username string) (*models.UserProfile, error) {
	var profile models.UserProfile
	if err := repo.db.Where("username = ?", username).First(&profile).Error; err != nil {
		return nil, err
	}
	return &profile, nil
}

func (repo *userProfileRepo) Delete(id uint) error {
	return repo.db.Delete(&models.UserProfile{}, id).Error
}
