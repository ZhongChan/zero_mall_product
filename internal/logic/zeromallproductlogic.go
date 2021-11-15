package logic

import (
	"context"

	"zero-mall/zero_mall_product/internal/svc"
	"zero-mall/zero_mall_product/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type Zero_mall_productLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewZero_mall_productLogic(ctx context.Context, svcCtx *svc.ServiceContext) Zero_mall_productLogic {
	return Zero_mall_productLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Zero_mall_productLogic) Zero_mall_product(req types.Request) (*types.Response, error) {
	// todo: add your logic here and delete this line

	return &types.Response{}, nil
}
