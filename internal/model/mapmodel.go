package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ MapModel = (*customMapModel)(nil)

type (
	// MapModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMapModel.
	MapModel interface {
		mapModel
	}

	customMapModel struct {
		*defaultMapModel
	}
)

// NewMapModel returns a model for the database table.
func NewMapModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) MapModel {
	return &customMapModel{
		defaultMapModel: newMapModel(conn, c, opts...),
	}
}
