package xfeed

import (
	"context"

	"github.com/pkg/errors"
	"github.com/x-feed/x-feed-sdk-golang/pkg/logger"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

type Client struct {
	conn *grpc.ClientConn

	Cfg Config

	// context and cancel func used to cancel all operations and gracefully stop client
	ctx    context.Context
	cancel context.CancelFunc

	Lg logger.LogEntry

	session *Session
}

func NewClient(cfg Config, logger logger.LogEntry) (*Client, error) {

	client := &Client{
		Cfg: cfg,
		Lg:  logger,
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

	client.Lg.Debugf("connection successful to host %s", cfg.ServerURI)

	go func() {
		<-client.ctx.Done()
		err := client.conn.Close()
		if err != nil {
			client.Lg.Errorf("connection close error %v", err)
		}
	}()

	client.session = &Session{
		clientConn:     client.conn,
		requestTimeout: cfg.RequestDeadline,
		Lg:             client.Lg,
		limiter:        rate.NewLimiter(rate.Limit(cfg.RequestRateLimit), cfg.RequestRateLimitBurst),
	}

	return client, nil
}

func (c *Client) GetSession() *Session {
	return c.session
}
