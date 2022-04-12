package server

import (
	"bufio"
	"log"
)

// clientSession implements the Session interface
type clientSession struct {
	id     int
	writer *bufio.Writer
	server *Server
}

func (session *clientSession) ID() int {
	return session.id
}

func (session *clientSession) SetID(id int) {
	session.id = id
	session.server.associateSession(id, session)
}

func (session *clientSession) Send(message []byte) {
	log.Printf("sending to message %d: %s", session.id, message)
	session.writer.Write(message)
	session.writer.WriteString("\n")
	session.writer.Flush()
}

func (session *clientSession) Subscribe(subscriptionID, id int) {
	session.server.subscribe(subscriptionID, id)
}

func (session *clientSession) Broadcast(id int, message []byte) []int {
	return session.server.sendBroadcast(id, message)
}
