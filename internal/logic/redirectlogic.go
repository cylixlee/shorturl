// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"database/sql"
	"errors"

	"github.com/cylixlee/shorturl/internal/svc"
	"github.com/cylixlee/shorturl/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var (
	ErrNotFound = errors.New("not found")
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
	u, err := l.svcCtx.MapModel.FindOneBySurl(l.ctx, sql.NullString{String: req.ShortURL, Valid: true})
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, ErrNotFound
		}
		logx.Errorw("error while finding map by surl", logx.Field("err", err))
		return nil, err
	}
	return &types.RedirectResponse{LongURL: u.Lurl.String}, nil
}
