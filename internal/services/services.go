package services

import (
	"context"
	"d0c/TestTaskBackDev/config"
	"d0c/TestTaskBackDev/helper"
	"d0c/TestTaskBackDev/internal/repo"
)

type Session interface {
	GenerateTokens(ctx context.Context, guid string) (Tokens, error)
	RefreshTokens(ctx context.Context, refreshToken string) (Tokens, error)
}

type Services struct {
	Session Session
}

type ServicesDependencies struct {
	SessionRepo repo.Session
	Hasher      helper.Hasher

	Config config.Config
}

func NewServices(deps ServicesDependencies) *Services {
	return &Services{
		Session: NewSessionService(deps.SessionRepo, deps.Hasher, deps.Config),
	}
}
