package server

import (
	"context"
	"git.xx.network/elixxir/mainnet-commitments/messages"
	"github.com/pkg/errors"
	"gitlab.com/xx_network/comms/connect"
	"gitlab.com/xx_network/primitives/id"
)

func StartServer(key, cert []byte, addr string) (*Impl, error) {
	pc, _, err := connect.StartCommServer(&id.Permissioning, addr, cert, key, nil)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to start comms server")
	}
	impl := &Impl{
		pc: pc,
	}
	messages.RegisterCommitmentsServer(pc.LocalServer, impl)
	return impl, nil
}

type Impl struct {
	pc *connect.ProtoComms
}

func (i *Impl) Verify(context.Context, *messages.Commitment) (*messages.Ack, error) {
	return nil, nil
}
