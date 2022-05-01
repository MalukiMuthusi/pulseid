package mysql

import (
	"context"
	"time"

	"github.com/MalukiMuthusi/pulseid/internal/logger"
	"github.com/MalukiMuthusi/pulseid/internal/models"
)

func (m MysqlDB) Active(ctx context.Context) ([]*models.Token, error) {

	db := m.Db

	db = db.Where("recalled = ?", false)

	db = db.Where("expiry < ", time.Now())

	var tokens []*models.Token

	res := db.Find(&tokens)

	if res.Error != nil {
		logger.Log.Info(res.Error)
		return nil, res.Error
	}

	return tokens, nil
}
