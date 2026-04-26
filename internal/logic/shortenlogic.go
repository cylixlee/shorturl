// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"net/url"
	"path"

	"github.com/open-portfolios/shorturl/internal/model"
	"github.com/open-portfolios/shorturl/internal/svc"
	"github.com/open-portfolios/shorturl/internal/types"
	"github.com/open-portfolios/shorturl/pkg/detect"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	ErrAlreadyShortURL = errors.New("this url is already shortened url")
)

type ShortenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShortenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShortenLogic {
	return &ShortenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShortenLogic) Shorten(req *types.ShortenRequest) (*types.ShortenResponse, error) {
	// Detect whether LongURL is reachable
	reachable, err := detect.Get(req.LongURL)
	if err != nil {
		logx.Errorw("unexpected error while detecting URL", logx.Field("err", err.Error()))
		return nil, err
	}
	if !reachable {
		logx.Errorf("URL %s unreachable", req.LongURL)
		return nil, err
	}

	// Check whether LongURL has been shortened before
	arr := md5.Sum([]byte(req.LongURL))
	sum := hex.EncodeToString(arr[:])
	v, err := l.svcCtx.MapModel.FindOneByMd5(l.ctx, sql.NullString{String: sum, Valid: true})
	if err == nil {
		return nil, fmt.Errorf("This URL has been shortened as %s", v.Surl.String)
	}
	if !errors.Is(err, model.ErrNotFound) {
		logx.Errorw("unexpected error while finding map by md5", logx.Field("err", err))
		return nil, err
	}

	// Check the LongURL is not actually a short url.
	u, err := url.Parse(req.LongURL)
	if err != nil {
		logx.Errorw("error parsing url", logx.Field("err", err))
		return nil, err
	}
	base := path.Base(u.Path)
	_, err = l.svcCtx.MapModel.FindOneBySurl(l.ctx, sql.NullString{String: base, Valid: true})
	if err == nil {
		return nil, ErrAlreadyShortURL
	}
	if !errors.Is(err, model.ErrNotFound) {
		logx.Errorw("unexpected error while finding map by surl", logx.Field("err", err))
		return nil, err
	}

	// Generate good short URL
	var short string
	for {
		id, err := l.svcCtx.Dispencer.Dispence(l.ctx)
		if err != nil {
			logx.Errorw("failed to dispence ID", logx.Field("err", err))
			return nil, err
		}
		short = string(l.svcCtx.Encoder.FormatUint(id))
		if l.svcCtx.Blacklist.Good(short) {
			break
		}
	}

	// Store short URL into database
	_, err = l.svcCtx.MapModel.Insert(l.ctx, &model.Map{
		Lurl: sql.NullString{String: req.LongURL, Valid: true},
		Md5:  sql.NullString{String: sum, Valid: true},
		Surl: sql.NullString{String: short, Valid: true},
	})
	if err != nil {
		logx.Errorw("error while inserting short URL into database", logx.Field("err", err))
		return nil, err
	}

	if err := l.svcCtx.BloomFilter.Add([]byte(short)); err != nil {
		logx.Errorw("error while adding short URL into BloomFilter", logx.Field("err", err))
		return nil, err
	}

	shortURL := path.Join(l.svcCtx.Config.ShortDomain, short)
	return &types.ShortenResponse{ShortURL: shortURL}, nil
}
