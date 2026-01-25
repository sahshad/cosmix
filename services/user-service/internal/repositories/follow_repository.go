package repositories

import (
	"user-service/internal/models"

	"gorm.io/gorm"
)

type FollowRepository interface {
	Create(follow *models.Follow) error
	Delete(followerID, followingID uint) error
	IsFollowing(followerID, followingID uint) (bool, error)
	GetFollowers(userID uint) ([]uint, error)
	GetFollowing(userID uint) ([]uint, error)
	GetFollowerCount(userID uint) (int64, error)
	GetFollowingCount(userID uint) (int64, error)
}

type followRepo struct {
	db *gorm.DB
}

func NewFollowRepository(db *gorm.DB) FollowRepository {
	return &followRepo{db: db}
}

func (repo *followRepo) Create(follow *models.Follow) error {
	return repo.db.Create(follow).Error
}

func (repo *followRepo) Delete(followerID, followingID uint) error {
	return repo.db.Where("follower_id = ? AND following_id = ?", followerID, followingID).
		Delete(&models.Follow{}).Error
}

func (repo *followRepo) IsFollowing(followerID, followingID uint) (bool, error) {
	var count int64
	err := repo.db.Model(&models.Follow{}).
		Where("follower_id = ? AND following_id = ?", followerID, followingID).
		Count(&count).Error
	return count > 0, err
}

func (repo *followRepo) GetFollowers(userID uint) ([]uint, error) {
	var followers []uint
	err := repo.db.Model(&models.Follow{}).
		Where("following_id = ?", userID).
		Pluck("follower_id", &followers).Error
	return followers, err
}

func (repo *followRepo) GetFollowing(userID uint) ([]uint, error) {
	var following []uint
	err := repo.db.Model(&models.Follow{}).
		Where("follower_id = ?", userID).
		Pluck("following_id", &following).Error
	return following, err
}

func (repo *followRepo) GetFollowerCount(userID uint) (int64, error) {
	var count int64
	err := repo.db.Model(&models.Follow{}).
		Where("following_id = ?", userID).
		Count(&count).Error
	return count, err
}

func (repo *followRepo) GetFollowingCount(userID uint) (int64, error) {
	var count int64
	err := repo.db.Model(&models.Follow{}).
		Where("follower_id = ?", userID).
		Count(&count).Error
	return count, err
}
