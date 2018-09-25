package xfeed

import (
	"time"
)

// Config represents configuration of x-feed Client
type Config struct {
	ClientID string

	// grpc transport specific parameters
	ServerURI string

	// DialTimeout is the timeout for failing to establish a connection.
	DialTimeout time.Duration

	RequestDeadline         time.Duration
	RequestRateLimit        float64
	RequestRateLimitBurst   int
	EntitiesRefreshInterval time.Duration
	MaxMessageSize          int

	// After a duration of this time if the client doesn't see any activity
	// it pings the server to see if the transport is still alive.
	// it is better to set it no more than 1 sec?
	InactiveTimeout time.Duration

	// After having pinged for keepalive check,
	// the client waits for a duration of Timeout and if no activity is seen even after that the connection is closed.
	KeepAliveTimeout time.Duration // The current default value is 20 seconds.

	// If true, client runs keepalive checks even with no active RPCs.
	PermitWithoutStream bool // false by default.

	//StatusChangeHandler stores function which is invoked any time when status is changed
	// that function accepts instance of Status struct instance
	StatusChangeHandler func(Status)
}
