// Code generated by goctl. DO NOT EDIT!
// Source: add.proto

package adder

import (
	"context"

	"zero_mall_product/rpc/add/add"

	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddReq  = add.AddReq
	AddResp = add.AddResp

	Adder interface {
		Add(ctx context.Context, in *AddReq, opts ...grpc.CallOption) (*AddResp, error)
	}

	defaultAdder struct {
		cli zrpc.Client
	}
)

func NewAdder(cli zrpc.Client) Adder {
	return &defaultAdder{
		cli: cli,
	}
}

func (m *defaultAdder) Add(ctx context.Context, in *AddReq, opts ...grpc.CallOption) (*AddResp, error) {
	client := add.NewAdderClient(m.cli.Conn())
	return client.Add(ctx, in, opts...)
}
