package mysql

import (
	"context"

	"github.com/MalukiMuthusi/pulseid/internal/logger"
	"github.com/MalukiMuthusi/pulseid/internal/models"
)

func (m MysqlDB) GetToken(ctx context.Context, tokenId string) (*models.Token, error) {
	var token models.Token

	tx := m.Db.Where("token = ?", tokenId).First(&token)

	if tx.Error != nil {
		logger.Log.Info(tx.Error)
		return nil, tx.Error
	}

	return &token, nil
}
