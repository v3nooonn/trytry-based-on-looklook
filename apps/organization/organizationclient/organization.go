// Code generated by goctl. DO NOT EDIT.
// Source: organization.proto

package organizationclient

import (
	"context"

	"github.com/v3nooonn/trytry/apps/organization/pb/organization"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	InfoReq  = organization.InfoReq
	InfoResp = organization.InfoResp

	Organization interface {
		Info(ctx context.Context, in *InfoReq, opts ...grpc.CallOption) (*InfoResp, error)
	}

	defaultOrganization struct {
		cli zrpc.Client
	}
)

func NewOrganization(cli zrpc.Client) Organization {
	return &defaultOrganization{
		cli: cli,
	}
}

func (m *defaultOrganization) Info(ctx context.Context, in *InfoReq, opts ...grpc.CallOption) (*InfoResp, error) {
	client := organization.NewOrganizationClient(m.cli.Conn())
	return client.Info(ctx, in, opts...)
}
