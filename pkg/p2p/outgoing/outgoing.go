package outgoing

import (
	"context"
	"net"
	"time"

	"github.com/pkg/errors"
	"github.com/wavesplatform/gowaves/pkg/libs/bytespool"
	"github.com/wavesplatform/gowaves/pkg/p2p/conn"
	"github.com/wavesplatform/gowaves/pkg/p2p/peer"
	"github.com/wavesplatform/gowaves/pkg/proto"
	"go.uber.org/zap"
)

type DuplicateChecker interface {
	Add([]byte) (isNew bool)
}

type EstablishParams struct {
	Address          proto.TCPAddr
	WavesNetwork     string
	Parent           peer.Parent
	Pool             bytespool.Pool
	DeclAddr         proto.TCPAddr
	Skip             conn.SkipFilter
	NodeName         string
	NodeNonce        uint64
	DuplicateChecker DuplicateChecker
}

func EstablishConnection(ctx context.Context, params EstablishParams, v proto.Version) error {
	ctx, cancel := context.WithCancel(ctx)
	remote := peer.NewRemote()
	p := connector{
		params: params,
		cancel: cancel,
		remote: remote,
	}

	c, err := net.Dial("tcp", params.Address.String())
	if err != nil {
		return err
	}

	connection, handshake, err := p.connect(ctx, c, v)
	if err != nil {
		zap.S().Debug(err, params.Address)
		return errors.Wrapf(err, "%q", params.Address)
	}
	p.connection = connection

	peerImpl := peer.NewPeerImpl(*handshake, connection, peer.Outgoing, remote, cancel)

	connected := peer.InfoMessage{
		Peer: peerImpl,
		Value: &peer.Connected{
			Peer: peerImpl,
		},
	}
	params.Parent.InfoCh <- connected
	zap.S().Debugf("connected %s, id: %s", params.Address, peerImpl.ID())

	return peer.Handle(peer.HandlerParams{
		Ctx:              ctx,
		ID:               peerImpl.ID(),
		Connection:       p.connection,
		Remote:           remote,
		Parent:           params.Parent,
		Pool:             params.Pool,
		Peer:             peerImpl,
		DuplicateChecker: params.DuplicateChecker,
	})
}

type connector struct {
	params     EstablishParams
	cancel     context.CancelFunc
	remote     peer.Remote
	connection conn.Connection
}

func (a *connector) connect(ctx context.Context, c net.Conn, v proto.Version) (conn.Connection, *proto.Handshake, error) {
	handshake := proto.Handshake{
		AppName:      a.params.WavesNetwork,
		Version:      v,
		NodeName:     a.params.NodeName,
		NodeNonce:    a.params.NodeNonce,
		DeclaredAddr: proto.HandshakeTCPAddr(a.params.DeclAddr),
		Timestamp:    proto.NewTimestampFromTime(time.Now()),
	}

	_, err := handshake.WriteTo(c)
	if err != nil {
		zap.S().Error("failed to send handshake: ", err, a.params.Address)
		return nil, nil, err
	}

	select {
	case <-ctx.Done():
		c.Close()
		return nil, nil, errors.Wrap(ctx.Err(), "connector.connect")
	default:
	}

	_, err = handshake.ReadFrom(c)
	if err != nil {
		zap.S().Debugf("failed to read handshake: %s %s", err, a.params.Address)
		select {
		case <-ctx.Done():
			return nil, nil, errors.Wrap(ctx.Err(), "connector.connect")
		case <-time.After(5 * time.Minute):
			return nil, nil, err
		}
	}
	return conn.WrapConnection(c, a.params.Pool, a.remote.ToCh, a.remote.FromCh, a.remote.ErrCh, a.params.Skip), &handshake, nil
}
