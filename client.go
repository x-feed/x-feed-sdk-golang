package xfeed

import (
	"context"

	"github.com/x-feed/x-feed-sdk-golang/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"sync"
)

type Client struct {
	conn *grpc.ClientConn

	cfg   Config
	creds *credentials.TransportCredentials

	// context and cancel func used to cancel all operations and gracefully stop client
	ctx    context.Context
	cancel context.CancelFunc

	lg logger.LogEntry

	m         sync.Mutex
	connected bool
	session   *session
}

func NewClient(cfg Config) (*Client, error) {
	return nil, nil
}
