package xa

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"context"
	"github.com/x-feed/x-feed-sdk-golang/pkg/logger"
)

type Client struct {
	conn *grpc.ClientConn

	cfg           Config
	creds         *credentials.TransportCredentials

	ctx    context.Context
	cancel context.CancelFunc

	lg *logger.LogEntry

}
