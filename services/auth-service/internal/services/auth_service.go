package services

import (
	"errors"

	"auth-service/internal/dto"
	"auth-service/internal/models"
	"auth-service/internal/repositories"
	"auth-service/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(input dto.RegisterDTO) (*models.User, error)
	Login(input dto.LoginDTO) (accessToken string, refreshToken string, user *models.User, err error)
	Refresh(refreshToken string) (newAccess string, newRefresh string, err error)
	GetByID(id uint) (*models.User, error)
}

type authService struct {
	userRepo repositories.UserRepository
}

func NewAuthService(userRepo repositories.UserRepository) AuthService {
	return &authService{userRepo: userRepo}
}

func (svc *authService) Register(input dto.RegisterDTO) (*models.User, error) {
	if _, err := svc.userRepo.FindByEmail(input.Email); err == nil {
		return nil, errors.New("email already in use")
	}
	if _, err := svc.userRepo.FindByUsername(input.Username); err == nil {
		return nil, errors.New("username already in use")
	}

	pwHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		FirstName:    input.FirstName,
		LastName:     input.LastName,
		Username:     input.Username,
		Email:        input.Email,
		PasswordHash: string(pwHash),
		DateOfBirth:  input.DateOfBirth,
		Role:         models.RoleUser,
	}

	if err := svc.userRepo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (svc *authService) Login(input dto.LoginDTO) (string, string, *models.User, error) {
	user, err := svc.userRepo.FindByEmail(input.Email)
	if err != nil {
		return "", "", nil, errors.New("invalid credentials")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)); err != nil {
		return "", "", nil, errors.New("invalid credentials")
	}

	access, err := utils.GenerateAccessToken(user.ID, string(user.Role))
	if err != nil {
		return "", "", nil, err
	}
	refresh, err := utils.GenerateRefreshToken(user.ID, string(user.Role))
	if err != nil {
		return "", "", nil, err
	}
	return access, refresh, user, nil
}

func (svc *authService) Refresh(refreshToken string) (string, string, error) {
	claims, err := utils.ParseRefreshToken(refreshToken)
	if err != nil {
		return "", "", errors.New("invalid refresh token")
	}

	if _, err := svc.userRepo.FindByID(claims.UserID); err != nil {
		return "", "", errors.New("invalid user")
	}

	newAccess, err := utils.GenerateAccessToken(claims.UserID, claims.Role)
	if err != nil {
		return "", "", err
	}
	newRefresh, err := utils.GenerateRefreshToken(claims.UserID, claims.Role)
	if err != nil {
		return "", "", err
	}
	return newAccess, newRefresh, nil
}


func (svc *authService) GetByID(id uint) (*models.User, error) {
	return svc.userRepo.FindByID(id)
}
