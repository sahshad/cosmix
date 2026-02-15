package repositories

import (
	"user-service/internal/models"

	"gorm.io/gorm"
)

type UserProfileRepositoryInterface interface {
	Create(profile *models.UserProfile) error
	FindByUserID(userID uint) (*models.UserProfile, error)
	FindByID(id uint) (*models.UserProfile, error)
	FindByUsername(username string) (*models.UserProfile, error)
	Update(profile *models.UserProfile) error
	Delete(id uint) error
}

type UserProfileRepo struct {
	db *gorm.DB
}

func NewUserProfileRepository(db *gorm.DB) *UserProfileRepo {
	return &UserProfileRepo{db: db}
}

func (repo *UserProfileRepo) Create(profile *models.UserProfile) error {
	return repo.db.Create(profile).Error
}

func (repo *UserProfileRepo) FindByUserID(userID uint) (*models.UserProfile, error) {
	var profile models.UserProfile
	if err := repo.db.Where("user_id = ?", userID).First(&profile).Error; err != nil {
		return nil, err
	}
	return &profile, nil
}

func (repo *UserProfileRepo) FindByID(id uint) (*models.UserProfile, error) {
	var profile models.UserProfile
	if err := repo.db.First(&profile, id).Error; err != nil {
		return nil, err
	}
	return &profile, nil
}

func (repo *UserProfileRepo) Update(profile *models.UserProfile) error {
	return repo.db.Save(profile).Error
}

func (repo *UserProfileRepo) FindByUsername(username string) (*models.UserProfile, error) {
	var profile models.UserProfile
	if err := repo.db.Where("username = ?", username).First(&profile).Error; err != nil {
		return nil, err
	}
	return &profile, nil
}

func (repo *UserProfileRepo) Delete(id uint) error {
	return repo.db.Delete(&models.UserProfile{}, id).Error
}
