package api

import (
	"context"
	"fmt"
	"log"

	"go-pre-interview/model"
	"go-pre-interview/server"
)

// API implements the server dispatch interface
type API struct {
	datastore *model.Datastore
}

// Init bootstraps the API
func Init(datastore *model.Datastore) *API {
	return &API{datastore}
}

type handlerFunc = func() error

// Dispatch executes a command in the given session
func (api *API) Dispatch(ctx context.Context, requestType string, requestID int, payload []byte) {

	session := ctx.Value(server.SessionKey).(Session)
	if session == nil {
		panic("no session in context")
	}
	err := api.dispatch(session, requestType, requestID, payload)
	if err != nil {
		session.Send(NewErrorResponse(requestID, err.Error()))
	}
}

func (api *API) dispatch(session Session, requestType string, requestID int, payload []byte) error {

	log.Printf("got api request %s %s", requestType, payload)

	dispatchTable := map[string]handlerFunc{
		"test.create":    func() error { return api.CreateHandler(session, requestID, payload) },
		"test.info":      func() error { return api.infoHandler(session, requestID, payload) },
		"test.broadcast": func() error { return api.broadcastHandler(session, requestID, payload) },
	}

	handler, found := dispatchTable[requestType]
	if !found {
		return fmt.Errorf("unimplemented API method %s", requestType)
	}

	switch requestType {
	case "test.create":
		break
	default:
		if session.ID() == 0 {
			return errCreatingTest
		}
	}

	return handler()
}
