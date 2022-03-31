package goclient

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/Jeffail/gabs/v2"
)

// Client ...
type Client struct {
	mutex          sync.Mutex
	requestID      uint32
	recursionDepth uint32
	clients        map[string]net.Conn
	outputs        map[string][]string
}

type responses struct {
	mutex sync.Mutex
	// map of requestID to response channel
	resps map[uint32]chan []byte
}

var (
	responseMap responses
)

func init() {
	responseMap = responses{
		resps: make(map[uint32]chan []byte),
	}
}

// New create a new client
func New() *Client {
	return &Client{
		clients: make(map[string]net.Conn),
		outputs: make(map[string][]string),
	}
}

// Connect ...
func (c *Client) Connect(username string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	connection, err := net.DialTimeout("tcp", ":8000", 1*time.Second)
	if err != nil {
		return err
	}
	log.Printf("got connection %+v", connection)
	c.clients[username] = connection

	// start a listener for each connection
	go c.listenForResponse(connection, username)
	return nil
}

func (c *Client) Disconnect(username string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	var connection, err = c.getConnection(username)
	if err != nil {
		return err
	}
	return connection.Close()
}

// SendRequest send a request to the server
func (c *Client) SendRequest(username, requestType string, data map[string]interface{}) (*gabs.Container, error) {

	var connection, err = c.getConnection(username)
	if err != nil {
		return nil, err
	}

	atomic.AddUint32(&c.requestID, 1)

	data["type"] = requestType
	data["request_id"] = c.requestID

	var bytes []byte
	bytes, err = json.Marshal(data)
	if err != nil {
		return nil, err
	}

	var delimiter = "\n"
	bytes = append(bytes, delimiter...)

	_, err = connection.Write(bytes)
	if err != nil {
		return nil, err
	}

	// set up receive channel for the reply
	var respChan = make(chan []byte)

	responseMap.mutex.Lock()
	responseMap.resps[c.requestID] = respChan
	responseMap.mutex.Unlock()

	for {
		select {
		case response := <-respChan:
			delete(responseMap.resps, c.requestID)
			return gabs.ParseJSON(response)
		case <-time.After(5 * time.Second):
			return nil, errors.New("timed_out")
		}
	}
}

func (c *Client) listenForResponse(connection net.Conn, username string) {
	// set SetReadDeadline
	err := connection.SetReadDeadline(time.Now().Add(5 * time.Second))
	if err != nil {
		log.Fatalln("SetReadDeadline failed:", err)
	}

	for {

		recvBuf := make([]byte, 2048)

		var n int
		n, err = connection.Read(recvBuf[:]) // recv data
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				// time out
				log.Fatalln("read timeout:", err)
			} else if err == io.EOF {
				err = c.addResultToRecives(recvBuf[0:n], username)
				if err != nil {
					log.Fatalln("received EOF and unable to add body to response map:", err)
				}
			} else if strings.Contains(err.Error(), "use of closed network connection") {
				// I know this is bad but go doesn't export this error, do nothing
				// plus go does this so there https://github.com/golang/go/blob/d422f54619b5b6e6301eaa3e9f22cfa7b65063c8/src/net/error_test.go#L507
			} else {
				// some error else, do something else, for example create new conn
				log.Fatalln("read error:", err)
			}
		}
		err = c.addResultToRecives(recvBuf[0:n], username)
		if err != nil {
			log.Fatalln("unable to add body to response map:", err)
		}

	}
}

func (c *Client) addResultToRecives(body []byte, username string) error {
	var messages = strings.Split(string(body), "\n")
	for _, m := range messages {
		if m == "" {
			// empty line skip
			continue
		}
		log.Printf("%s received message: %s\n", username, m)
		var (
			err       error
			cleanBody = []byte(strings.TrimSpace(m))
			j         *gabs.Container
		)
		j, err = gabs.ParseJSON(cleanBody)
		if err != nil {
			log.Println("error parsing json:", err)
			return err
		}
		if !j.ExistsP("reply_to") {
			// no one is listing for this message add outputs and continue
			c.mutex.Lock()
			c.outputs[username] = append(c.outputs[username], string(cleanBody))
			c.mutex.Unlock()
			continue
		}
		var replyTo = uint32(j.Path("reply_to").Data().(float64))

		responseMap.resps[replyTo] <- cleanBody
	}
	return nil
}

func (c *Client) getConnection(username string) (net.Conn, error) {
	var connection net.Conn
	var ok bool
	if connection, ok = c.clients[username]; !ok {
		return nil, errors.New("unknown user " + username)
	}
	return connection, nil
}
