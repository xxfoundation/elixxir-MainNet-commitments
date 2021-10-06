package client

import (
	"git.xx.network/elixxir/mainnet-commitments/messages"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/pkg/errors"
	"gitlab.com/xx_network/comms/connect"
	"gitlab.com/xx_network/crypto/signature/rsa"
	"gitlab.com/xx_network/primitives/id"
	"google.golang.org/grpc"
)

type Sender interface {
	TransmitSignature(host *connect.Host, message *messages.Commitment) error
}

// Client struct implements the GRPC client call to mainnet-commitments servers
type Client struct {
	pc *connect.ProtoComms
}

// StartClient func creates a client
func StartClient(key, salt []byte, id *id.ID) (*Client, error) {
	pk, err := rsa.LoadPrivateKeyFromPem(key)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to load key")
	}
	pc, err := connect.CreateCommClient(id, rsa.CreatePublicKeyPem(pk.GetPublic()), key, salt)
	if err != nil {
		return nil, err
	}
	return &Client{
		pc: pc,
	}, nil
}

// TransmitSignature func sends a Commitment message to the mainnet-commitments server
func (c *Client) TransmitSignature(host *connect.Host, message *messages.Commitment) error {
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
