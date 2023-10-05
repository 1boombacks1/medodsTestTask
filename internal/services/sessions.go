package services

import (
	"context"
	"d0c/TestTaskBackDev/config"
	"d0c/TestTaskBackDev/helper"
	"d0c/TestTaskBackDev/internal/models"
	"d0c/TestTaskBackDev/internal/repo"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

type SessionService struct {
	sessionRepo repo.Session
	hasher      helper.Hasher

	SignKey         string
	JWTokenTTL      time.Duration
	RefreshTokenTTL time.Duration
}

func NewSessionService(sessionRepo repo.Session, hasher helper.Hasher, cfg config.Config) *SessionService {
	return &SessionService{
		sessionRepo: sessionRepo,
		hasher:      hasher,

		SignKey:         cfg.SignKey,
		JWTokenTTL:      cfg.JWTokenTTL,
		RefreshTokenTTL: cfg.RefreshTokenTTL,
	}
}

type Tokens struct {
	AccessToken     string
	RefreshToken    string
	RefreshTokenTTL int
}

func (s *SessionService) GenerateTokens(ctx context.Context, guid string) (Tokens, error) {
	const operation = "services.SessionService.GenerateTokens"

	claims := helper.AuthClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(s.JWTokenTTL)),
			IssuedAt:  jwt.Now(),
		},
		Guid: guid,
	}

	accessToken, err := helper.MakeJWT(claims, s.SignKey)
	if err != nil {
		return Tokens{}, fmt.Errorf("%s - %w", operation, err)
	}

	refreshToken := s.hasher.Hash(helper.NewRefreshToken())

	modelSession := models.Session{
		Guid:         guid,
		RefreshToken: refreshToken,
		ExpiresIn:    time.Now().Add(s.RefreshTokenTTL).Unix(),
	}

	if _, err = s.sessionRepo.CreateSession(ctx, &modelSession); err != nil {
		if errors.Is(err, repo.ErrAlreadyExists) {
			s.sessionRepo.DropSession(ctx, &modelSession)
			s.sessionRepo.CreateSession(ctx, &modelSession)
		} else {
			return Tokens{}, fmt.Errorf("%s - s.sessionRepo.CreateSession: %w", operation, err)
		}
	}

	tokens := Tokens{
		AccessToken:     accessToken,
		RefreshToken:    base64.StdEncoding.EncodeToString([]byte(refreshToken)),
		RefreshTokenTTL: int(s.RefreshTokenTTL.Seconds()),
	}

	return tokens, nil
}

func (s *SessionService) RefreshTokens(ctx context.Context, encodedRefreshToken string) (Tokens, error) {
	const operation = "services.SessionService.RefreshTokens"

	decodedRefreshToken, err := base64.StdEncoding.DecodeString(encodedRefreshToken)
	if err != nil {
		return Tokens{}, fmt.Errorf("%s - base64.StdEncoding.DecodeString: %w", operation, err)
	}

	session, err := s.sessionRepo.GetSessionByRefreshToken(ctx, string(decodedRefreshToken))
	if err != nil {
		if errors.Is(err, repo.ErrNotFound) {
			return Tokens{}, ErrInvalidRefreshSession
		}
		return Tokens{}, fmt.Errorf("%s - s.sessionRepo.GetSessionByRefreshToken: %w", operation, err)
	}

	s.sessionRepo.DropSession(ctx, session)

	if time.Now().Unix() >= session.ExpiresIn {
		return Tokens{}, ErrRefreshTokenExpired
	}

	claims := helper.AuthClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(s.JWTokenTTL)),
			IssuedAt:  jwt.Now(),
		},
		Guid: session.Guid,
	}
	accessToken, err := helper.MakeJWT(claims, s.SignKey)
	if err != nil {
		return Tokens{}, fmt.Errorf("%s - %w", operation, err)
	}

	refreshToken := s.hasher.Hash(helper.NewRefreshToken())

	session = &models.Session{
		RefreshToken: refreshToken,
		ExpiresIn:    time.Now().Add(s.RefreshTokenTTL).Unix(),
	}

	if _, err = s.sessionRepo.CreateSession(ctx, session); err != nil {
		return Tokens{}, fmt.Errorf("%s - s.sessionRepo.CreateSession: %w", operation, err)
	}

	tokens := Tokens{
		AccessToken:     accessToken,
		RefreshToken:    base64.StdEncoding.EncodeToString([]byte(refreshToken)),
		RefreshTokenTTL: int(s.RefreshTokenTTL.Seconds()),
	}

	return tokens, nil
}
