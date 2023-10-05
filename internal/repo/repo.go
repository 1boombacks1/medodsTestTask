package repo

import (
	"context"
	"d0c/TestTaskBackDev/internal/models"
)

//go:generate moq --out=mock/mock_repo.go --pkg=mock . Session
type Session interface {
	CreateSession(ctx context.Context, session *models.Session) (string, error)
	DropSession(ctx context.Context, session *models.Session) error
	GetSessionByRefreshToken(ctx context.Context, refreshToken string) (*models.Session, error)
}
