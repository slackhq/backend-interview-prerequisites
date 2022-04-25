package server

import (
	"bufio"
	"context"
	"encoding/json"
	"io"
	"log"
	"net"
	"sync"
)

const (
	listenAddr = "localhost:8000"
)

// Server wraps the server side state
type Server struct {
	// global lock
	lock sync.Mutex

	subscriptions map[int][]*clientSession

	sessions map[int][]*clientSession
}

// StartServer creates and starts new server
func StartServer(dispatcher Dispatcher) {
	var s = Server{
		subscriptions: map[int][]*clientSession{},
		sessions:      map[int][]*clientSession{},
	}
	s.listen(dispatcher)
}

// listen runs forever accepting connections on the socket and dispatching
// incoming requests
func (s *Server) listen(dispatcher Dispatcher) {
	l, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatal("listen error:", err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		}

		go s.serve(conn, dispatcher)
	}
}

// serve runs in a goroutine for each connection and parses incoming messages
func (s *Server) serve(conn net.Conn, dispatcher Dispatcher) {
	log.Printf("got new connection")

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	session := clientSession{
		writer: writer,
		server: s,
	}

	for {
		payload, err := reader.ReadBytes('\n')
		if err != nil {
			if err != io.EOF {
				log.Printf("error reading from socket: %v", err)
			}
			log.Printf("session disconnected %d", session.ID())
			s.lock.Lock()
			defer s.lock.Unlock()

			s.unsubscribeSession(&session)
			return
		}

		s.dispatch(dispatcher, &session, payload)
	}
}

// dispatch processes a single incoming message
func (s *Server) dispatch(dispatcher Dispatcher, session *clientSession, payload []byte) {
	var request map[string]interface{}
	err := json.Unmarshal(payload, &request)
	if err != nil {
		log.Printf("error decoding request: %v", err)
		return
	}

	requestType, ok := request["type"]
	if !ok {
		log.Printf("error decoding request: no type")
		return
	}

	requestIDVar, ok := request["request_id"]
	if !ok {
		log.Printf("error decoding request: no request ID")
		return
	}

	// json decodes numerics as float by default
	requestID := int(requestIDVar.(float64))

	// simple concurrency model -- lock the server for each dispatch
	s.lock.Lock()
	defer s.lock.Unlock()

	dispatcher.Dispatch(context.WithValue(context.Background(), SessionKey, session), requestType.(string), requestID, payload)
}

func (s *Server) associateSession(id int, session *clientSession) {
	sessions, _ := s.sessions[id]
	if sessions == nil {
		sessions = make([]*clientSession, 0, 10)
	}
	s.sessions[id] = append(sessions, session)
}

func (s *Server) subscribe(subscriptionID, id int) {
	sessions, _ := s.sessions[id]
	if sessions == nil {
		return
	}

	for _, session := range sessions {
		subscriptions, _ := s.subscriptions[subscriptionID]
		if subscriptions == nil {
			subscriptions = make([]*clientSession, 0, 10)
		}

		s.subscriptions[subscriptionID] = append(subscriptions, session)
	}
}

func (s *Server) sendBroadcast(subscriptionID int, message []byte) []int {
	sessions, found := s.subscriptions[subscriptionID]
	ids := make([]int, 0, 10)
	if !found {
		return ids
	}
	for _, session := range sessions {
		session.Send(message)
		ids = append(ids, session.ID())
	}
	return ids
}

// unsubscribeSession removes all current subscriptions for the given session
func (s *Server) unsubscribeSession(deadSession *clientSession) {
	for id, sessions := range s.sessions {
		// Filter in place adapted from https://github.com/golang/go/wiki/SliceTricks
		n := 0
		for _, session := range sessions {
			if session != deadSession {
				sessions[n] = session
				n++
			}
		}
		s.sessions[id] = sessions[:n]
	}

	for subscriptionID, sessions := range s.subscriptions {
		// Filter in place adapted from https://github.com/golang/go/wiki/SliceTricks
		n := 0
		for _, session := range sessions {
			if session != deadSession {
				sessions[n] = session
				n++
			}
		}
		s.subscriptions[subscriptionID] = sessions[:n]
	}
}
