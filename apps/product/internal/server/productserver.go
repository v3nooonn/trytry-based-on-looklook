// Code generated by goctl. DO NOT EDIT.
// Source: product.proto

package server

import (
	"context"

	"github.com/v3nooonn/trytry/apps/product/internal/logic"
	"github.com/v3nooonn/trytry/apps/product/internal/svc"
	"github.com/v3nooonn/trytry/apps/product/pb/product"
)

type ProductServer struct {
	svcCtx *svc.ServiceContext
	product.UnimplementedProductServer
}

func NewProductServer(svcCtx *svc.ServiceContext) *ProductServer {
	return &ProductServer{
		svcCtx: svcCtx,
	}
}

func (s *ProductServer) Pagination(ctx context.Context, in *product.PaginationReq) (*product.PaginationResp, error) {
	l := logic.NewPaginationLogic(ctx, s.svcCtx)
	return l.Pagination(in)
}

func (s *ProductServer) Info(ctx context.Context, in *product.InfoReq) (*product.InfoResp, error) {
	l := logic.NewInfoLogic(ctx, s.svcCtx)
	return l.Info(in)
}
