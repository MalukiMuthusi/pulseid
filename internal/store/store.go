package store

import (
	"context"

	"github.com/MalukiMuthusi/pulseid/internal/models"
)

type Store interface {
	SaveToken(ctx context.Context, token *models.Token) (*models.Token, error)
}

type MockStore struct {
}

func NewMockStore() MockStore {
	return MockStore{}
}

func (m MockStore) SaveToken(ctx context.Context, token *models.Token) (*models.Token, error) {

	return token, nil
}
