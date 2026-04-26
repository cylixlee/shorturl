// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package handler

import (
	"errors"
	"net/http"

	"github.com/open-portfolios/shorturl/internal/logic"
	"github.com/open-portfolios/shorturl/internal/model"
	"github.com/open-portfolios/shorturl/internal/svc"
	"github.com/open-portfolios/shorturl/internal/types"
	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RedirectHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RedirectRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		if err := validator.New().StructCtx(r.Context(), &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewRedirectLogic(r.Context(), svcCtx)
		resp, err := l.Redirect(&req)
		if err != nil {
			if errors.Is(err, model.ErrNotFound) {
				http.NotFound(w, r)
				return
			}
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		http.Redirect(w, r, resp.LongURL, http.StatusFound)
	}
}
