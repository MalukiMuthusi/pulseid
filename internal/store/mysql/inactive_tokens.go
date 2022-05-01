package mysql

import (
	"context"
	"time"

	"github.com/MalukiMuthusi/pulseid/internal/logger"
	"github.com/MalukiMuthusi/pulseid/internal/models"
)

func (m MysqlDB) Inactive(ctx context.Context) ([]*models.Token, error) {

	var tokens []*models.Token

	res := m.Db.Where("recalled = ?", true).Or("expiry > ?", time.Now()).Find(&tokens)

	if res.Error != nil {
		logger.Log.Info(res.Error)
		return nil, res.Error
	}

	return tokens, nil
}
