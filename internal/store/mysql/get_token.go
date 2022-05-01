package mysql

import (
	"context"

	"github.com/MalukiMuthusi/pulseid/internal/models"
)

func (m MysqlDB) GetToken(ctx context.Context, tokenId string) (*models.Token, error) {
	panic("not implemented")
}
