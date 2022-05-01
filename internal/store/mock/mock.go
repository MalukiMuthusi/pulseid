package mock

import (
	"context"
	"math/rand"
	"time"

	"github.com/MalukiMuthusi/pulseid/internal/models"
)

type Store struct {
}

func NewStore() Store {
	return Store{}
}

func (m Store) SaveToken(ctx context.Context, token *models.Token) (*models.Token, error) {

	return token, nil
}

func (m Store) GetToken(ctx context.Context, tokenId string) (*models.Token, error) {

	token, err := models.NewToken()

	if err != nil {
		return nil, err
	}

	return token, nil
}

func (m Store) RecallToken(ctx context.Context, tokenId string) (*models.Token, error) {
	token, err := models.NewToken()
	if err != nil {
		return nil, err
	}

	token.Recalled = true

	return token, nil
}

func (m Store) Active(ctx context.Context) ([]*models.Token, error) {

	var tokens []*models.Token

	n := rand.Intn(14)

	for i := 0; i < n; i++ {

		token, err := models.NewToken()

		token.Recalled = false

		token.CreatedAt = time.Now().AddDate(0, 0, -4)

		if err != nil {
			return nil, err
		}

		tokens = append(tokens, token)
	}

	return tokens, nil
}

func (m Store) Inactive(ctx context.Context) ([]*models.Token, error) {
	var tokens []*models.Token

	n := rand.Intn(14)

	for i := 0; i < n; i++ {

		token, err := models.NewToken()

		token.Recalled = true

		if err != nil {
			return nil, err
		}

		tokens = append(tokens, token)
	}

	return tokens, nil
}
