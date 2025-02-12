package service

import (
	"strconv"

	"github.com/0bvim/octoBuddy/internal/domain/entity"
	"github.com/0bvim/octoBuddy/internal/domain/repository"
	"github.com/0bvim/octoBuddy/internal/infrastructure/auth/github"
	"github.com/0bvim/octoBuddy/internal/infrastructure/auth/jwt"
)

type AuthService struct {
	githubClient *github.Client
	tokenService *jwt.TokenService
	userRepo     repository.UserRepository
}

func NewAuthService(
	githubClient *github.Client,
	tokenService *jwt.TokenService,
	userRepo repository.UserRepository,
) *AuthService {
	return &AuthService{
		githubClient: githubClient,
		tokenService: tokenService,
		userRepo:     userRepo,
	}
}

func (s *AuthService) GetAuthURL() string {
	return s.githubClient.GetAuthCodeURL("state-token")
}

func (s *AuthService) HandleCallback(code string) (*entity.TokenPair, *entity.User, error) {
	// Exchange code for token
	token, err := s.githubClient.ExchangeCode(code)
	if err != nil {
		return nil, nil, err
	}

	// Get user info from GitHub
	githubUser, err := s.githubClient.GetUserInfo(token)
	if err != nil {
		return nil, nil, err
	}

	// Save or update user in repository
	existingUser, _ := s.userRepo.FindByID(strconv.Itoa(githubUser.ID))
	if existingUser == nil {
		if err := s.userRepo.Save(githubUser); err != nil {
			return nil, nil, err
		}
	}

	// Generate JWT tokens
	tokenPair, err := s.tokenService.GenerateTokenPair(githubUser)
	if err != nil {
		return nil, nil, err
	}

	return tokenPair, githubUser, nil
}

func (s *AuthService) RefreshToken(refreshToken string) (*entity.TokenPair, error) {
	claims, err := s.tokenService.ParseToken(refreshToken)
	if err != nil {
		return nil, err
	}

	user, err := s.userRepo.FindByID(claims.UserID)
	if err != nil {
		return nil, err
	}

	return s.tokenService.GenerateTokenPair(user)
}
