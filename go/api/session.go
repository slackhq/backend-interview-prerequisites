package api

// Session wraps a user session
type Session interface {
	// UserID returns the user id for the current session
	ID() int

	// SetUserID sets the user id in the current session (to be used in login)
	SetID(userID int)

	// Send a reply on the current connection
	Send(message []byte)

	// Subscribe the given user ID to the given channel ID
	Subscribe(channelID, userID int)

	// Broadcast the message to all sessions subscribed to the given channel.
	// Returns the set of connected users that received the message.
	Broadcast(channelID int, message []byte) []int
}
