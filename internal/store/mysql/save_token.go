package mysql

import (
	"context"

	"github.com/MalukiMuthusi/pulseid/internal/logger"
	"github.com/MalukiMuthusi/pulseid/internal/models"
)

func (m MysqlDB) SaveToken(ctx context.Context, token *models.Token) (*models.Token, error) {

	tx := m.Db.Create(token)

	if tx.Error != nil {
		logger.Log.Info(tx.Error)
		return nil, tx.Error
	}

	return token, nil
}
