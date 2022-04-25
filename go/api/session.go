package api

// Session wraps a session
type Session interface {
	// ID returns the id for the current session
	ID() int

	// SetID sets the id in the current session
	SetID(id int)

	// Send a reply on the current connection
	Send(message []byte)

	// Subscribe the given ID
	Subscribe(sessionid, id int)

	// Broadcast a message.
	Broadcast(id int, message []byte) []int
}
