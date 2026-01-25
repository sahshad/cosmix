package services

import (
	"errors"
	"user-service/internal/models"
	"user-service/internal/repositories"
)

type FollowService interface {
	Follow(followerID uint, followingID uint) error
	Unfollow(followerID uint, followingID uint) error
	GetFollowers(userID uint) ([]uint, error)
	GetFollowing(userID uint) ([]uint, error)
}

type followService struct {
	repo repositories.FollowRepository
}

func NewFollowService(repo repositories.FollowRepository) FollowService {
	return &followService{
		repo: repo,
	}
}

func (svc *followService) Follow(followerID, followingID uint) error {
	if followerID == followingID {
		return errors.New("cannot follow yourself")
	}

	isFollowing, err := svc.repo.IsFollowing(followerID, followingID)
	if err != nil {
		return err
	}
	if isFollowing {
		return errors.New("already following")
	}

	follow := &models.Follow{
		FollowerID:  followerID,
		FollowingID: followingID,
	}
	return svc.repo.Create(follow)
}

func (svc *followService) Unfollow(followerID, followingID uint) error {
	return svc.repo.Delete(followerID, followingID)
}

func (svc *followService) GetFollowers(userID uint) ([]uint, error) {
	return svc.repo.GetFollowers(userID)
}

func (svc *followService) GetFollowing(userID uint) ([]uint, error) {
	return svc.repo.GetFollowing(userID)
}
