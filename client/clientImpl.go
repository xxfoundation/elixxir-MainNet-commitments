package client

import (
	"git.xx.network/elixxir/mainnet-commitments/messages"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/pkg/errors"
	"gitlab.com/xx_network/comms/connect"
	"gitlab.com/xx_network/primitives/id"
	"google.golang.org/grpc"
)

type Sender interface {
	SignAndTransmit(host *connect.Host, message *messages.Commitment) error
}

// Client struct implements the GRPC client call to mainnet-commitments servers
type Client struct {
	pc *connect.ProtoComms
}

// StartClient func creates a client
func StartClient(key, cert, salt []byte, id *id.ID) (*Client, error) {
	pc, err := connect.CreateCommClient(id, cert, key, salt)
	if err != nil {
		return nil, err
	}
	return &Client{
		pc: pc,
	}, nil
}

// SignAndTransmit func sends a Commitment message to the mainnet-commitments server
func (c *Client) SignAndTransmit(host *connect.Host, message *messages.Commitment) error {
	f := func(conn *grpc.ClientConn) (*any.Any, error) {
		// Set up the context
		ctx, cancel := host.GetMessagingContext()
		defer cancel()

		// Send the message
		resultMsg, err := messages.NewCommitmentsClient(conn).Verify(ctx, message)
		if err != nil {
			err = errors.New(err.Error())
			return nil, errors.New(err.Error())

		}
		return ptypes.MarshalAny(resultMsg)
	}
	_, err := c.pc.Send(host, f)
	return err
}
