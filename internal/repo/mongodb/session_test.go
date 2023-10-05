package mongodb_test

import (
	"context"
	"d0c/TestTaskBackDev/internal/models"
	"testing"
	"time"
)

func TestCreateSession(t *testing.T) {
	type args struct {
		ctx     context.Context
		session models.Session
	}

	testCases := []struct {
		name         string
		args         args
		mockBehavior func(context.Context, *models.Session) (string, error)
		want         string
		wantErr      bool
	}{
		{
			name: "OK",
			args: args{
				ctx: context.Background(),
				session: models.Session{
					Guid:      "db654f5d2c354b118cccaaa1fb6cc81f",
					ExpiresIn: time.Now().Add(5 * time.Minute).Unix(),
				},
			},
			mockBehavior: func(ctx context.Context, s *models.Session) (string, error) {
				return "123", nil
			},
			want:    "123",
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// sessionRepo := mongodb.NewSessionRepository()
		})
	}
}
