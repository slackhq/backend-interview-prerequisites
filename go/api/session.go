package api

// Session wraps a user session
type Session interface {
	// ID returns the id for the current session
	ID() int

	// SetID sets the id in the current session
	SetID(userID int)

	// Send a reply on the current connection
	Send(message []byte)

	// Subscribe the given ID
	Subscribe(id, userID int)

	// Broadcast a message.
	Broadcast(id int, message []byte) []int
}
