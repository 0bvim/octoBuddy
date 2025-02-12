package jwt

import (
	"strconv"
	"time"

	"github.com/0bvim/octoBuddy/internal/domain/entity"
	"github.com/golang-jwt/jwt/v4"
)

type TokenService struct {
	secret string
}

func NewTokenService(secret string) *TokenService {
	return &TokenService{secret: secret}
}

// internal/infrastructure/auth/jwt/token_service.go

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func (s *TokenService) GenerateTokenPair(user *entity.User) (*entity.TokenPair, error) {
	// Access token (15 minutes expiration)
	accessTokenClaims := &Claims{
		UserID: strconv.Itoa(user.ID),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessToken.SignedString([]byte(s.secret))
	if err != nil {
		return nil, err
	}

	// Refresh token (7 days expiration)
	refreshTokenClaims := &Claims{
		UserID: strconv.Itoa(user.ID),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * 7 * time.Hour)),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(s.secret))
	if err != nil {
		return nil, err
	}

	return &entity.TokenPair{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}

func (s *TokenService) ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}
