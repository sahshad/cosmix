package repositories

import (
	"auth-service/internal/models"

	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	Create(u *models.User) error
	FindByEmail(email string) (*models.User, error)
	FindByID(id uint) (*models.User, error)
	FindByUsername(username string) (*models.User, error)
	Update(u *models.User) error
}

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (repo *UserRepo) Create(u *models.User) error {
	return repo.db.Create(u).Error
}

func (repo *UserRepo) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := repo.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepo) FindByID(id uint) (*models.User, error) {
	var user models.User
	if err := repo.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepo) FindByUsername(username string) (*models.User, error) {
	var user models.User
	if err := repo.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepo) Update(u *models.User) error {
	return repo.db.Save(u).Error
}
