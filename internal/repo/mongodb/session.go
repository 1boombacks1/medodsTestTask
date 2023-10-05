package mongodb

import (
	"context"
	"d0c/TestTaskBackDev/internal/models"
	"d0c/TestTaskBackDev/internal/repo"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SessionRepository struct {
	collection *mongo.Collection
}

func NewSessionRepository(db *mongo.Database) *SessionRepository {
	return &SessionRepository{
		collection: db.Collection("sessions"),
	}
}

func (s *SessionRepository) CreateSession(ctx context.Context, session *models.Session) (string, error) {
	const operation = "mongodb.SessionStorage.CreateSession"

	session.CreatedAt = time.Now().UTC()
	model := fromModelToMongo(session)

	count, err := s.collection.CountDocuments(ctx, bson.M{"_id": model.Id})
	if err != nil {
		return "", fmt.Errorf("%s - s.collection.CountDocuments: %w", operation, err)
	}

	if count > 0 {
		return "", repo.ErrAlreadyExists
	}

	res, err := s.collection.InsertOne(ctx, model)
	if err != nil {
		return "", fmt.Errorf("%s - s.collection.InsertOne: %w", operation, err)
	}

	id := res.InsertedID.(primitive.ObjectID).String()

	return id, nil
}

func (s *SessionRepository) GetSessionByRefreshToken(ctx context.Context, refreshToken string) (*models.Session, error) {
	const operation = "mongodb.SessionStorage.GetSessionByRefreshToken"

	session := new(SessionMongo)

	if err := s.collection.FindOne(ctx, bson.M{"refresh_token": refreshToken}).Decode(&session); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, repo.ErrNotFound
		}
		return nil, fmt.Errorf("%s - s.collection.FindOne: %w", operation, err)
	}

	return fromMongoToModel(session), nil
}

func (s *SessionRepository) DropSession(ctx context.Context, session *models.Session) error {
	const operation = "mongodb.SessionStorage.DropSession"

	model := fromModelToMongo(session)

	if _, err := s.collection.DeleteOne(ctx, bson.M{"guid": model.Guid}); err != nil {
		return fmt.Errorf("%s - s.db.DeleteOne: %w", operation, err)
	}

	return nil
}

type SessionMongo struct {
	Id           string    `bson:"_id,omitempty"`
	Guid         string    `bson:"guid"`
	RefreshToken string    `bson:"refresh_token"`
	ExpiresIn    int64     `bson:"expires_in"`
	CreatedAt    time.Time `bson:"created_at"`
}

func fromMongoToModel(session *SessionMongo) *models.Session {
	return &models.Session{
		Id:           session.Id,
		Guid:         session.Guid,
		RefreshToken: session.RefreshToken,
		ExpiresIn:    session.ExpiresIn,
		CreatedAt:    session.CreatedAt,
	}
}

func fromModelToMongo(session *models.Session) *SessionMongo {
	return &SessionMongo{
		Id:           session.Id,
		Guid:         session.Guid,
		RefreshToken: session.RefreshToken,
		ExpiresIn:    session.ExpiresIn,
		CreatedAt:    session.CreatedAt,
	}
}
