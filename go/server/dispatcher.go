package server

import "context"

type contextKey int

const (
	// SessionKey defines where in the Context the session is stored
	SessionKey contextKey = iota
)

// Dispatcher is the interface used by the server to dispatch incoming requests
type Dispatcher interface {
	Dispatch(ctx context.Context, requestType string, requestID int, payload []byte)
}
