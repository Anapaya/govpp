// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package policer

import (
	"context"
	"fmt"
	"io"

	api "go.fd.io/govpp/api"
	memclnt "go.fd.io/govpp/binapi/memclnt"
)

// RPCService defines RPC service policer.
type RPCService interface {
	PolicerAdd(ctx context.Context, in *PolicerAdd) (*PolicerAddReply, error)
	PolicerAddDel(ctx context.Context, in *PolicerAddDel) (*PolicerAddDelReply, error)
	PolicerBind(ctx context.Context, in *PolicerBind) (*PolicerBindReply, error)
	PolicerBindV2(ctx context.Context, in *PolicerBindV2) (*PolicerBindV2Reply, error)
	PolicerDel(ctx context.Context, in *PolicerDel) (*PolicerDelReply, error)
	PolicerDump(ctx context.Context, in *PolicerDump) (RPCService_PolicerDumpClient, error)
	PolicerDumpV2(ctx context.Context, in *PolicerDumpV2) (RPCService_PolicerDumpV2Client, error)
	PolicerInput(ctx context.Context, in *PolicerInput) (*PolicerInputReply, error)
	PolicerInputV2(ctx context.Context, in *PolicerInputV2) (*PolicerInputV2Reply, error)
	PolicerOutput(ctx context.Context, in *PolicerOutput) (*PolicerOutputReply, error)
	PolicerOutputV2(ctx context.Context, in *PolicerOutputV2) (*PolicerOutputV2Reply, error)
	PolicerReset(ctx context.Context, in *PolicerReset) (*PolicerResetReply, error)
	PolicerUpdate(ctx context.Context, in *PolicerUpdate) (*PolicerUpdateReply, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) PolicerAdd(ctx context.Context, in *PolicerAdd) (*PolicerAddReply, error) {
	out := new(PolicerAddReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) PolicerAddDel(ctx context.Context, in *PolicerAddDel) (*PolicerAddDelReply, error) {
	out := new(PolicerAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) PolicerBind(ctx context.Context, in *PolicerBind) (*PolicerBindReply, error) {
	out := new(PolicerBindReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) PolicerBindV2(ctx context.Context, in *PolicerBindV2) (*PolicerBindV2Reply, error) {
	out := new(PolicerBindV2Reply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) PolicerDel(ctx context.Context, in *PolicerDel) (*PolicerDelReply, error) {
	out := new(PolicerDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) PolicerDump(ctx context.Context, in *PolicerDump) (RPCService_PolicerDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_PolicerDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_PolicerDumpClient interface {
	Recv() (*PolicerDetails, error)
	api.Stream
}

type serviceClient_PolicerDumpClient struct {
	api.Stream
}

func (c *serviceClient_PolicerDumpClient) Recv() (*PolicerDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *PolicerDetails:
		return m, nil
	case *memclnt.ControlPingReply:
		err = c.Stream.Close()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) PolicerDumpV2(ctx context.Context, in *PolicerDumpV2) (RPCService_PolicerDumpV2Client, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_PolicerDumpV2Client{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_PolicerDumpV2Client interface {
	Recv() (*PolicerDetails, error)
	api.Stream
}

type serviceClient_PolicerDumpV2Client struct {
	api.Stream
}

func (c *serviceClient_PolicerDumpV2Client) Recv() (*PolicerDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *PolicerDetails:
		return m, nil
	case *memclnt.ControlPingReply:
		err = c.Stream.Close()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) PolicerInput(ctx context.Context, in *PolicerInput) (*PolicerInputReply, error) {
	out := new(PolicerInputReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) PolicerInputV2(ctx context.Context, in *PolicerInputV2) (*PolicerInputV2Reply, error) {
	out := new(PolicerInputV2Reply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) PolicerOutput(ctx context.Context, in *PolicerOutput) (*PolicerOutputReply, error) {
	out := new(PolicerOutputReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) PolicerOutputV2(ctx context.Context, in *PolicerOutputV2) (*PolicerOutputV2Reply, error) {
	out := new(PolicerOutputV2Reply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) PolicerReset(ctx context.Context, in *PolicerReset) (*PolicerResetReply, error) {
	out := new(PolicerResetReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) PolicerUpdate(ctx context.Context, in *PolicerUpdate) (*PolicerUpdateReply, error) {
	out := new(PolicerUpdateReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}
