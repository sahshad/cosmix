package services

import (
	"errors"
	"user-service/internal/models"
	"user-service/internal/repositories"
)

type FollowServiceInterface interface {
	Follow(followerID uint, followingID uint) error
	Unfollow(followerID uint, followingID uint) error
	GetFollowers(userID uint) ([]uint, error)
	GetFollowing(userID uint) ([]uint, error)
}

type FollowService struct {
	repo repositories.FollowRepositoryInterface
}

func NewFollowService(repo repositories.FollowRepositoryInterface) FollowServiceInterface {
	return &FollowService{
		repo: repo,
	}
}

func (svc *FollowService) Follow(followerID, followingID uint) error {
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

func (svc *FollowService) Unfollow(followerID, followingID uint) error {
	return svc.repo.Delete(followerID, followingID)
}

func (svc *FollowService) GetFollowers(userID uint) ([]uint, error) {
	return svc.repo.GetFollowers(userID)
}

func (svc *FollowService) GetFollowing(userID uint) ([]uint, error) {
	return svc.repo.GetFollowing(userID)
}
