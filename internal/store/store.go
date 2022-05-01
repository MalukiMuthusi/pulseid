package store

import (
	"context"
	"time"

	"github.com/MalukiMuthusi/pulseid/internal/models"
	"gorm.io/gorm"
)

type Store interface {
	SaveToken(ctx context.Context, token *models.Token) (*models.Token, error)
	GetToken(ctx context.Context, tokenId string) (*models.Token, error)
	RecallToken(ctx context.Context, tokenId string) (*models.Token, error)
}

type MockStore struct {
}

func NewMockStore() MockStore {
	return MockStore{}
}

func (m MockStore) SaveToken(ctx context.Context, token *models.Token) (*models.Token, error) {

	return token, nil
}

func (m MockStore) GetToken(ctx context.Context, tokenId string) (*models.Token, error) {
	return &models.Token{
		Model: gorm.Model{
			CreatedAt: time.Now().AddDate(1, 0, 0),
		},
	}, nil
}

func (m MockStore) RecallToken(ctx context.Context, tokenId string) (*models.Token, error) {
	return &models.Token{
		Recalled: true,
	}, nil
}
