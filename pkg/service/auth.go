package service

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	"github.com/Inexpediency/todo-rest-api/pkg/dto"
	"github.com/Inexpediency/todo-rest-api/pkg/models"
	"github.com/Inexpediency/todo-rest-api/pkg/repository"
)

const (
	accessTokenTTL        = time.Hour * 2
	refreshTokenTTL       = time.Hour * 24 * 30
	accessTokenSigningKey = "81hJ!*@#Y&12yN#UI!Yjfklsjdf"
	refreshTokenSigningKey = "410fj12fjhsdfjksaj(UY^JIJ98adsuJIKDiHA&*"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user models.User) (int, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	user.Password = string(hashedPassword)

	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateTokens(signInDto dto.SignIn) (dto.Tokens, error) {
	var tokens dto.Tokens

	user, err := s.repo.GetUserByUsername(signInDto.Username)
	if err != nil {
		return tokens, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(signInDto.Password)); err != nil {
		return tokens, err
	}

	accessTokenRaw := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(accessTokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: user.Id,
	})
	accessToken, err := accessTokenRaw.SignedString([]byte(accessTokenSigningKey))
	if err != nil {
		return tokens, err
	}

	refreshTokenRaw := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(refreshTokenTTL).Unix(),
			IssuedAt: time.Now().Unix(),
		},
		UserId: user.Id,
	})
	refreshToken, err := refreshTokenRaw.SignedString([]byte(refreshTokenSigningKey))
	if err != nil {
		return tokens, err
	}
	if err := s.repo.SaveRefreshToken(user.Id, refreshToken); err != nil {
		return tokens, err
	}

	tokens.AccessToken = accessToken
	tokens.RefreshToken = refreshToken

	return tokens, nil
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(accessTokenSigningKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}
