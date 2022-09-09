// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package mss_clamp

import (
	"context"
	"fmt"
	"io"

	api "go.fd.io/govpp/api"
)

// RPCService defines RPC service mss_clamp.
type RPCService interface {
	MssClampEnableDisable(ctx context.Context, in *MssClampEnableDisable) (*MssClampEnableDisableReply, error)
	MssClampGet(ctx context.Context, in *MssClampGet) (RPCService_MssClampGetClient, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) MssClampEnableDisable(ctx context.Context, in *MssClampEnableDisable) (*MssClampEnableDisableReply, error) {
	out := new(MssClampEnableDisableReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) MssClampGet(ctx context.Context, in *MssClampGet) (RPCService_MssClampGetClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_MssClampGetClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_MssClampGetClient interface {
	Recv() (*MssClampDetails, error)
	api.Stream
}

type serviceClient_MssClampGetClient struct {
	api.Stream
}

func (c *serviceClient_MssClampGetClient) Recv() (*MssClampDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *MssClampDetails:
		return m, nil
	case *MssClampGetReply:
		err = c.Stream.Close()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}