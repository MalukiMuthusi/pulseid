package store

import (
	"context"

	"github.com/MalukiMuthusi/pulseid/internal/models"
)

type Store interface {
	SaveToken(ctx context.Context, token *models.Token) (*models.Token, error)

	GetToken(ctx context.Context, tokenId string) (*models.Token, error)

	RecallToken(ctx context.Context, tokenId string) (*models.Token, error)

	Active(ctx context.Context) ([]*models.Token, error)

	Inactive(ctx context.Context) ([]*models.Token, error)
}
