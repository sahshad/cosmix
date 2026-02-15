package services

import (
	"errors"
	"time"
	"user-service/internal/dto"
	"user-service/internal/models"
	"user-service/internal/repositories"
)

type UserProfileServiceInterface interface {
	GetProfile(userID uint) (*dto.UserProfileResponse, error)
	GetProfileByID(id uint) (*dto.UserProfileResponse, error)
	GetProfileByUsername(username string) (*dto.UserProfileResponse, error)
	UpdateProfile(userID uint, input dto.UpdateProfileDTO) (*dto.UserProfileResponse, error)
	CreateProfile(profile *models.UserProfile) error
	CreateFromAuthEvent(event dto.UserCreatedFromDTO) error
}

type UserProfileService struct {
	repo repositories.UserProfileRepositoryInterface
}

func NewUserProfileService(repo repositories.UserProfileRepositoryInterface) UserProfileServiceInterface {
	return &UserProfileService{
		repo: repo,
	}
}

func (svc *UserProfileService) GetProfile(userID uint) (*dto.UserProfileResponse, error) {
	profile, err := svc.repo.FindByUserID(userID)
	if err != nil {
		return nil, errors.New("profile not found")
	}
	return svc.toResponse(profile), nil
}

func (svc *UserProfileService) GetProfileByID(id uint) (*dto.UserProfileResponse, error) {
	profile, err := svc.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("profile not found")
	}
	return svc.toResponse(profile), nil
}

func (svc *UserProfileService) UpdateProfile(userID uint, input dto.UpdateProfileDTO) (*dto.UserProfileResponse, error) {
	profile, err := svc.repo.FindByUserID(userID)
	if err != nil {
		return nil, errors.New("profile not found")
	}

	if input.FirstName != nil {
		profile.FirstName = *input.FirstName
	}
	if input.LastName != nil {
		profile.LastName = *input.LastName
	}
	if input.Username != nil {
		profile.Username = input.Username
	}
	if input.DateOfBirth != nil {
		dob, err := time.Parse("2006-01-02", *input.DateOfBirth)
		if err != nil {
			return nil, errors.New("invalid date of birth format")
		}
		profile.DateOfBirth = &dob
	}
	if input.AvatarURL != nil {
		profile.AvatarURL = input.AvatarURL
	}
	if input.Bio != nil {
		profile.Bio = input.Bio
	}

	if err := svc.repo.Update(profile); err != nil {
		return nil, err
	}

	return svc.toResponse(profile), nil
}

func (svc *UserProfileService) CreateProfile(profile *models.UserProfile) error {
	return svc.repo.Create(profile)
}

func (svc *UserProfileService) CreateFromAuthEvent(event dto.UserCreatedFromDTO) error {
	profile := &models.UserProfile{
		UserID:    event.AuthUserID,
		Email:     event.Email,
		FirstName: event.FirstName,
		LastName:  event.LastName,
		CreatedAt: event.CreatedAt,
	}
	return svc.repo.Create(profile)
}

func (svc *UserProfileService) GetProfileByUsername(username string) (*dto.UserProfileResponse, error) {
	profile, err := svc.repo.FindByUsername(username)
	if err != nil {
		return nil, errors.New("profile not found")
	}
	return svc.toResponse(profile), nil
}

func (svc *UserProfileService) toResponse(profile *models.UserProfile) *dto.UserProfileResponse {
	return &dto.UserProfileResponse{
		ID:          profile.ID,
		UserID:      profile.UserID,
		FirstName:   profile.FirstName,
		LastName:    profile.LastName,
		Username:    profile.Username,
		DateOfBirth: profile.DateOfBirth,
		AvatarURL:   profile.AvatarURL,
		Bio:         profile.Bio,
		CreatedAt:   profile.CreatedAt,
		UpdatedAt:   profile.UpdatedAt,
	}
}
