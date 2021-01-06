package service

import (
	"errors"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	"github.com/Inexpediency/todo-rest-api/pkg/dto"
	"github.com/Inexpediency/todo-rest-api/pkg/models"
	"github.com/Inexpediency/todo-rest-api/pkg/repository"
)

const (
	accessTokenTTL         = time.Hour * 2
	refreshTokenTTL        = time.Hour * 24 * 30
	accessTokenSigningKey  = "81hJ!*@#Y&12yN#UI!Yjfklsjdf"
	refreshTokenSigningKey = "410fj12fjhsdfjksaj(UY^JIJ98adsuJIKDiHA&*"
)

const (
	ACCESS_TOKEN  = iota
	REFRESH_TOKEN = iota
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
	user, err := s.repo.GetUserByUsername(signInDto.Username)
	if err != nil {
		return dto.Tokens{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(signInDto.Password)); err != nil {
		return dto.Tokens{}, err
	}

	return s.generateTokensFromId(user.Id)
}

func (s *AuthService) ParseToken(token string, tokenType int) (int, error) {
	var signingKey string

	switch tokenType {
	case ACCESS_TOKEN:
		signingKey = accessTokenSigningKey
		break
	case REFRESH_TOKEN:
		signingKey = refreshTokenSigningKey
		break
	default:
		return 0, errors.New("invalid token type")
	}

	parsedToken, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := parsedToken.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func (s *AuthService) RefreshTokens(refreshToken string) (dto.Tokens, error) {
	userId, err := s.ParseToken(refreshToken, REFRESH_TOKEN)
	if err != nil {
		return dto.Tokens{}, errors.New("invalid refresh token")
	}

	savedRefreshToken, err := s.repo.GetRefreshToken(userId)
	if err != nil {
		return dto.Tokens{}, errors.New("there is no refresh token for you; try to sign in")
	}

	if strings.Compare(savedRefreshToken, refreshToken) != 0 {
		return dto.Tokens{}, errors.New("invalid refresh token")
	}

	return s.generateTokensFromId(userId)
}

func (s *AuthService) generateTokensFromId(userId int) (dto.Tokens, error) {
	accessTokenRaw := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(accessTokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: userId,
	})
	accessToken, err := accessTokenRaw.SignedString([]byte(accessTokenSigningKey))
	if err != nil {
		return dto.Tokens{}, err
	}

	refreshTokenRaw := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(refreshTokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: userId,
	})

	refreshToken, err := refreshTokenRaw.SignedString([]byte(refreshTokenSigningKey))
	if err != nil {
		return dto.Tokens{}, err
	}
	if err := s.repo.SaveRefreshToken(userId, refreshToken); err != nil {
		return dto.Tokens{}, err
	}

	return dto.Tokens{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}
