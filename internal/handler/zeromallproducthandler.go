package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"zero-mall/zero_mall_product/internal/logic"
	"zero-mall/zero_mall_product/internal/svc"
	"zero-mall/zero_mall_product/internal/types"
)

func Zero_mall_productHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewZero_mall_productLogic(r.Context(), ctx)
		resp, err := l.Zero_mall_product(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
