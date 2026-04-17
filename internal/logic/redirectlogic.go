// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"database/sql"
	"errors"

	"github.com/cylixlee/shorturl/internal/model"
	"github.com/cylixlee/shorturl/internal/svc"
	"github.com/cylixlee/shorturl/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RedirectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRedirectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RedirectLogic {
	return &RedirectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RedirectLogic) Redirect(req *types.RedirectRequest) (*types.RedirectResponse, error) {
	exists, err := l.svcCtx.BloomFilter.ExistsCtx(l.ctx, []byte(req.ShortURL))
	if err != nil {
		logx.Errorw("error while querying BloomFilter", logx.Field("err", err))
		return nil, err
	}
	if !exists {
		return nil, model.ErrNotFound
	}

	u, err := l.svcCtx.MapModel.FindOneBySurl(l.ctx, sql.NullString{String: req.ShortURL, Valid: true})
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, err
		}
		logx.Errorw("error while finding map by surl", logx.Field("err", err))
		return nil, err
	}
	return &types.RedirectResponse{LongURL: u.Lurl.String}, nil
}
