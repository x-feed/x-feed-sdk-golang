package xfeed

import (
	"context"

	"github.com/pkg/errors"
	"github.com/x-feed/x-feed-sdk-golang/pkg/logging"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

// Client represents X-feed client
type Client struct {
	conn *grpc.ClientConn

	cfg Config

	// context and cancel func used to cancel all operations and gracefully stop client
	ctx    context.Context
	cancel context.CancelFunc

	logger logging.Logger

	session *Session
}

// NewClient provides casual way for creating the Client instance
func NewClient(cfg Config, logger logging.Logger) (*Client, error) {

	client := &Client{
		cfg:    cfg,
		logger: logger,
	}

	keepaliveCfg := keepalive.ClientParameters{
		Time:                cfg.InactiveTimeout,
		Timeout:             cfg.KeepAliveTimeout,
		PermitWithoutStream: cfg.PermitWithoutStream,
	}

	opts := []grpc.DialOption{
		grpc.WithInsecure(), //TODO: discuss with x-feed team and fix security
		grpc.WithKeepaliveParams(keepaliveCfg),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(cfg.MaxMessageSize),
			grpc.MaxCallSendMsgSize(cfg.MaxMessageSize),
		),
	}

	client.ctx, client.cancel = context.WithCancel(context.Background())

	var err error

	client.conn, err = grpc.DialContext(client.ctx, cfg.ServerURI, opts...)
	if err != nil {
		return nil, errors.Errorf("grpc dial err: %v", err)
	}

	client.logger.Debugf("connection successful to host %s", cfg.ServerURI)

	go func() {
		<-client.ctx.Done()
		err := client.conn.Close()
		if err != nil {
			client.logger.Errorf("connection close error %v", err)
		}
	}()

	client.session = &Session{
		clientID:       cfg.ClientID,
		clientConn:     client.conn,
		requestTimeout: cfg.RequestDeadline,
		logger:         client.logger,
		limiter:        rate.NewLimiter(rate.Limit(cfg.RequestRateLimit), cfg.RequestRateLimitBurst),
	}

	return client, nil
}

// Session returns instance of session in case where grpc connection is ready
func (c *Client) Session() (*Session, error) {
	if c == nil {
		return nil, errors.New("client is not initialized")
	}

	return c.session, nil
}
